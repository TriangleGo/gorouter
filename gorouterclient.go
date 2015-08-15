package main

import (
	"bufio"
	"fmt"
	"gorouter/network/protocol"
	"gorouter/network/simplebuffer"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

var reBase = regexp.MustCompile(`base +href="(.*?)"`)
var reHTML = regexp.MustCompile(`\saction=["']?(.*?)["'\s]|\shref=["']?(.*?)["'\s]|\ssrc=["']?(.*?)["'\s]|\surl\(["']?(.*?)["']?\)`)
var reCSS = regexp.MustCompile(`url\(["']?(.*?)["']?\)`)
var uri *url.URL
var urlPrefix = "http://127.0.0.1:8080/foo?url="
var domainPrefix = "http://127.0.0.1:8080"

func HTTPServer() {
	http.HandleFunc("/foo", fooHandler)
	http.ListenAndServe(":8080", nil)
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	reqUrl := r.FormValue("url")
	if reqUrl == "" {
		return
	}

	//处理域名重定向问题
	if strings.Contains(reqUrl, "www.") {
		if !strings.Contains(reqUrl, ".www.") {
			reqUrl = strings.Replace(reqUrl, "www.", "", 1)
		}
	}

	fmt.Printf(">>URL : %v \n", reqUrl)
	httpClient := &http.Client{}
	req, err := http.NewRequest(r.Method, reqUrl, nil)
	if err != nil {
		fmt.Printf(">>>[Error]:%v \n", err)
	}
	req.Header = r.Header
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Printf(">>>[Error]:%v \n", err)
	}

	for k, v := range resp.Header {
		for i := 0; i < len(v); i++ {
			w.Header().Add(k, v[i])
		}
	}
	w.WriteHeader(resp.StatusCode)

	//3xx redirection
	if resp.StatusCode > 300 && resp.StatusCode < 399 {
		fmt.Printf("redirection trigger %v \nResponse %v \n", resp.StatusCode, resp)
	}

	contentType := resp.Header.Get("Content-Type")
	// Rewrite all urls
	if strings.Contains(contentType, "text/html") {
		reWriteHtml(resp, r.URL.String())
	} else if strings.Contains(contentType, "text/css") {
		reWriteCss(resp, r.URL.String())
	} else {
		io.Copy(w, resp.Body)
	}

}

func reWriteHtml(resp *http.Response, urlRequest string) []byte {
	body, _ := ioutil.ReadAll(resp.Body)
	// if there's a <base href> specified in the document
	// use that as base to encode all URLs in the page
	baseHrefMatch := reBase.FindSubmatch(body)
	if len(baseHrefMatch) > 0 {
		var err error
		uri, err = url.Parse(string(baseHrefMatch[1][:]))
		fmt.Printf("uri %v \n", uri)
		if err != nil {
			log.Println("Error Parsing " + string(baseHrefMatch[1][:]))
		}
	}

	encodedBody := reHTML.ReplaceAllFunc(body, func(s []byte) []byte {
		parts := reHTML.FindSubmatchIndex(s)

		if parts != nil {
			// replace src attribute
			srcIndex := parts[2:4]
			if srcIndex[0] != -1 {
				return rewriteURI(s, srcIndex[0], srcIndex[1], urlRequest)
			}

			// replace href attribute
			hrefIndex := parts[4:6]
			if hrefIndex[0] != -1 {
				return rewriteURI(s, hrefIndex[0], hrefIndex[1], urlRequest)
			}

			// replace form action attribute
			actionIndex := parts[6:8]
			if actionIndex[0] != -1 {
				return rewriteURI(s, actionIndex[0], actionIndex[1], urlRequest)
			}

			// replace form url attribute
			urlIndex := parts[8:10]
			if urlIndex[0] != -1 {
				return rewriteURI(s, urlIndex[0], urlIndex[1], urlRequest)
			}
		}
		return s
	})
	return encodedBody
}

func reWriteCss(resp *http.Response, urlRequest string) []byte {
	body, _ := ioutil.ReadAll(resp.Body)
	encodedBody := reCSS.ReplaceAllFunc(body, func(s []byte) []byte {
		parts := reCSS.FindSubmatchIndex(s)
		if parts != nil {
			// replace url attribute in css
			pathIndex := parts[2:4]
			if pathIndex[0] != -1 {
				return rewriteURI(s, pathIndex[0], pathIndex[1], urlRequest)
			}
		}
		return s
	})
	return encodedBody
}

func rewriteURI(src []byte, start int, end int, fixUri string) []byte {
	relURL := string(src[start:end])
	// keep anchor and javascript links intact
	if relURL == "" || strings.HasPrefix(relURL, "#") || strings.HasPrefix(relURL, "javascript") || strings.HasPrefix(relURL, "data") {
		return src
	}
	// Check if url is relative and make it absolute
	if strings.Index(relURL, "http") != 0 {
		relPath, err := url.Parse(relURL)
		if err != nil {
			return src
		}

		relPath.String()

	}

	siteUrl := string(src[start:end])
	fmt.Printf("Parse URL %v\n", string(src[start:end]))
	if strings.Contains(siteUrl, `//`) {
		Pos := strings.Index(string(src), `http://`)
		if Pos < 1 {
			return src
		}
		return []byte(string(src)[0:Pos] + urlPrefix + string(src)[Pos:])
	} else if siteUrl[0:1] == `/` {
		//URL补全
		Pos := strings.Index(string(src), `/`)
		if Pos < 1 {
			return src
		}
		return []byte(string(src)[0:Pos] + domainPrefix + fixUri + string(src)[Pos:])
	} else {
		//
		return src
	}

	return src
}

func main() {
	fmt.Print("Go Router Client Runing \n")
	connClient, err := net.Dial("tcp", "127.0.0.1:9090")
	if err != nil {
		fmt.Print("Client Connecting error \n")
		return
	}

	buffer := simplebuffer.NewSimpleBuffer("bigEndian")
	buffer.WriteUInt32(10)
	buffer.WriteUInt8(4)
	buffer.WriteUInt8(3)
	buffer.WriteData([]byte("1234"))
	fmt.Printf("send data %v \n", buffer.Data())

	connClient.Write(buffer.Data())

	for {
		fmt.Print("Recv Looping  \n")
		buf := make([]byte, 4096)
		n, err := connClient.Read(buf)
		if err != nil {
			fmt.Printf("Client Read Buffer Failed %v %v\r\n", err, n)
			break
		}

		proto := protocol.NewProtocal()
		_, err = proto.PraseFromData(buf[0:n], n)
		if err != nil {
			fmt.Printf("Data Parse failed %v\r\n", err)
			continue
		}

		//fmt.Printf("Recv Data : %v \n\n", string(proto.Data))

		go func(p *protocol.Protocol) {
			reader := strings.NewReader(string(proto.Data))
			bufreader := bufio.NewReader(reader)
			req, err := http.ReadRequest(bufreader)

			if err != nil {
				fmt.Printf("construct request error %v \n", err)
				return
			}
			req.Host = "192.168.1.26:8306"
			url, err := url.Parse(`http://192.168.1.26:8306` + req.URL.Path)
			if err != nil {
				fmt.Printf("Parse Url Error %v \n", err)
			}
			req.URL = url
			req.RequestURI = ""
			fmt.Printf("construct request  %v \n", req)

			httpClient := &http.Client{}
			resp, err := httpClient.Do(req)
			if err != nil {
				fmt.Printf("client do request error %v \n", err)
				return
			}

			fmt.Printf("\n\n\n\n%v %v\nResponse \n%v \n", req.Method, req.URL, resp)

			//resp, _ = http.Get(`http://192.168.1.26:8306/abc.txt`)
			//fmt.Printf("Normal Get %v \n", resp)

		}(proto)
	}
}
