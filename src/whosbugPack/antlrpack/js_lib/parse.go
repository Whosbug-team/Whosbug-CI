package parser

import (
	"bytes"
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Node struct {
	Name  string
	Data  bytes.Buffer
	Attrs map[string]string
	// offsets to replace in file i.e someBytes[StartOffest:endOffest+1]
	StartOffest int
	EndOffest   int
}

// GetNodes pulls raw content for tags passed. If getAll is true nested are returned else just the root level.
func GetNodes(rd io.Reader, getAll bool, tags ...string) ([]*Node, error) {

	tkr := html.NewTokenizer(rd)

	var nd *Node

	allNodes := []*Node{}
	// where in file we are
	offset := 0
	// shows at root or nested
	depth := 0
	// track when we have a script and style
	var inTagWeWant = map[string]bool{}
	for _, t := range tags {
		inTagWeWant[strings.ToLower(t)] = true

	}
	var current = ""
	for {
		tt := tkr.Next()
		// content and keep track of offset
		rw := tkr.Raw()
		offset += len(rw)
		switch tt {
		case html.ErrorToken:
			// EOF normally...
			if err := tkr.Err(); err != nil && err != io.EOF {
				return nil, err
			}
			return allNodes, nil
		case html.TextToken:
			// if we want all occurence or just root level
			if (getAll || depth == 1) && inTagWeWant[current] {
				// if in a tag we want just write content
				if _, err := nd.Data.Write(rw); err != nil {
					return allNodes, err
				}

			}
		case html.StartTagToken:
			tagn, has := tkr.TagName()
			current = string(tagn)
			if (getAll || depth == 0) && inTagWeWant[current] {

				nd = &Node{Attrs: map[string]string{}}
				// any attrs
				for has {
					k, v, hasMore := tkr.TagAttr()
					nd.Attrs[string(k)] = string(v)
					has = hasMore

				}
				nd.StartOffest = offset - len(rw)
				if _, err := nd.Data.Write(rw); err != nil {
					return allNodes, err
				}

			}
			depth += 1
		case html.EndTagToken:
			tagn, _ := tkr.TagName()

			if (getAll || depth == 1) && inTagWeWant[string(tagn)] {
				nd.EndOffest = offset
				nd.Name = string(tagn)
				if _, err := nd.Data.Write(rw); err != nil {
					return allNodes, err
				}

				allNodes = append(allNodes, nd)

			}
			// set to empty or text nodes will get added at root
			current = ""
			depth -= 1

		}
	}

}

// Just the content and offsets
func JustOneContent(rd io.Reader, tag []byte) (bytes.Buffer, int, int, error) {

	// offset to replace in file i.e someBytes[startOffest:endOffest+1]
	startOffest := 0
	offset := 0
	depth := 0
	tkr := html.NewTokenizer(rd)
	var data bytes.Buffer
	// track when we are in tag we want
	var inTag = false
	for {
		tt := tkr.Next()
		// content and keep track of offset
		rw := tkr.Raw()
		// always inc end
		offset += len(rw)
		switch tt {
		case html.ErrorToken:
			// EOF normally...
			if err := tkr.Err(); err != nil && err != io.EOF {
				return data, startOffest, offset, err
			}
			return data, startOffest, offset, tkr.Err()
		case html.TextToken:
			// if in a script just write content
			if inTag && depth == 1 {
				if _, err := data.Write(rw); err != nil {
					return data, startOffest, offset, err
				}
			}
		case html.StartTagToken:
			tagn, _ := tkr.TagName()
			depth++
			if depth == 1 && bytes.Equal(tagn, tag) {
				startOffest = offset - len(rw)
				inTag = true
				startOffest -= len(rw)
				if _, err := data.Write(rw); err != nil {
					return data, startOffest, offset, err
				}
			}
		case html.EndTagToken:
			tagn, _ := tkr.TagName()
			depth--
			if tt == html.EndTagToken {
				// end tag for script so bail
				if depth == 0 && bytes.Equal(tagn, tag) {
					if _, err := data.Write(rw); err != nil {
						return data, startOffest, offset, err
					}
					return data, startOffest, offset, nil

				}

			}

		}
	}

}
