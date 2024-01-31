# SmuggleFuzz
HTTP/2 based downgrade and smuggle scanner

For more info see [https://moopinger.github.io/blog/smugglefuzz/fuzzing/smuggling/2024/01/31/SmuggleFuzz.html](https://moopinger.github.io/blog/smugglefuzz/fuzzing/smuggling/2024/01/31/SmuggleFuzz.html)


# Wordlist

SmuggleFuzz allows users complete control over requests through custom wordlists. These lists have a basic structure, which should be followed for optimal request handling. For instance, headers and their values are split using a `semicolon and a space (; )` instead of the usual colon, facilitating the inclusion of colon values in smuggling requests. This also opens up possibilities for various mutations and creative approaches. For detailed guidance on creating your own payloads, refer to: PortSwigger's HTTP/2 Research.

The tool includes a ready-to-use list of 125 smuggling gadgets, though there's always scope for expansion. These gadgets can be displayed using the 'output' command, providing insights into query structuring. Users can run scans with custom wordlists using the 'w' flag. The list supports URL encoding (%00) for non-printable byte values, such as carriage return and line feed represented as '%0d%0a' or '\r\n'. While the provided list is comprehensive, crafting your own gadgets can significantly enhance success rates.



### Pseudo headers


Pseudo headers are fully supported and can be customized using the ":name" syntax:

* `:authority`
* `:scheme`
* `:method`

These headers are particularly useful when testing for H2.0 attack vectors. For example, in testing H2.0 smuggling using OAST-based methods, wordlist entries like the following can be effective:

* `:authority; localhost\r\n\r\nGET / HTTP/1.1\r\nHost: uniqid.oastify.com\r\nX-HEADER:%20`

or

* `:authority; localhost\r\n\r\nGET https://uniqid.oastify.com/ HTTP/1.1\r\nHost: localhost\r\nX-HEADER:%20`

# Detection

The detection method employed in SmuggleFuzz is akin to those used by Albinowax in his HTTP/1 smuggling research and defparam's Smuggler, albeit with adaptations for HTTP/2. Fortunately, as these attacks involve downgrades, the same request body/DATA payload can be utilized for both H2.CL and H2.TE detection.

A crucial guideline for wordlists is that successfully smuggled headers should include the following values for a confirmatory attack:

    CL: 13
    TE: chunked

Each request incorporates a "Body"/DATA-Frame with the value "68". If a TE header is smuggled, it is interpreted as a chunked content length of 104 (0x68) by the server if successfully smuggled, leading to a timeout. Similarly, smuggling a CL header with a content length of 13 also results in a timeout due to unfulfilled length.

The confirmatory request resubmits the same query that led to the timeout, but with Data: "3\r\nABC\r\n0\r\n\r\n". This satisfies both a chunked TE and a CL of 13, eliciting a successful response and potentially indicating a smuggling vulnerability.

# Install

```
go install github.com/moopinger/smugglefuzz@latest
```


Compile yourself:


```
git clone https://github.com/moopinger/smugglefuzz.git
cd smugglefuzz
go build .
./smugglefuzz
```

# Usage

There are three commands for SmuggleFuzz:

* scan

```

Usage:
  smugglefuzz scan [flags]

Examples:
	smugglefuzz scan -u https://example.com/ --confirm
	smugglefuzz scan -u https://example.com/ --filter 200 --confirm
	smugglefuzz scan -u https://example.com/ -w wordlist.txt -t 10 --confirm
	smugglefuzz scan -dc -u https://example.com/ -w wordlist.txt -x PUT --confirm

Multiple Targets:
	smugglefuzz scan -f multiple_targets.txt --confirm -t 10
	smugglefuzz scan -f multiple_targets.txt -w wordlist.txt --confirm -s ./save-success.txt
	smugglefuzz scan -f multiple_targets.txt -w wordlist.txt -H "Cookie: date=...; session=...;" --confirm -s ./save-success.txt -x PUT

Flags:
  -c, --confirm               Enable this flag to send a confirmation to the target when a timeout is encountered. Helps confirm if the target is vulnerable.
      --dc                    Disable colour in the output. This is useful when you want to save the output to a file.
  -f, --file string           A file containing multiple targets in url format. One target per line.
      --filter string         Filter responses by string or frame type, etc. For example: 405, 200, 502, TIMEOUT, RST, GOAWAY, etc.
  -H, --header string         Insert a custom header. It should be provided in the regular header format: "Cookie: date=...; session=...;"
  -h, --help                  help for scan
  -i, --interval int          The timeout interval in seconds. (default 5)
  -x, --method string         The HTTP request method to be used. (default "POST")
  -s, --save-success string   If a request is confirmed to be successful (via the --confirm flag), it will be saved to a file. This is useful when dealing with lots of targets.
  -t, --threads int           The number of threads to run. Smugglefuzz can go fast, so set the desired number. However, too many may upset any WAFs. (default 4)
  -u, --url string            The target URL to be scanned.
  -w, --wordlist string       Provide a custom list of gadgets to use. If not provided, the default list will be used.

```



* request

```

  smugglefuzz request [flags]

Examples:
	smugglefuzz request -u https://www.example.com/ -a "content-length\t; 13"
	smugglefuzz request -u https://www.example.com/ -a "content-length\t; 13" -x PUT
	smugglefuzz request -u https://www.example.com/ -a "content-length\t; 13" -H "Cookie: date=...; session=...;"

Flags:
  -a, --attack string   Attack Header, separated by (; ) like the wordlist in 'scan' mode.
      --dc              Disable colour in the output.
  -H, --header string   Insert custom header. eg "Cookie: values"
  -h, --help            help for request
  -i, --interval int    Detection timeout interval in seconds. (default 5)
  -x, --method string   The method to use. (default "POST")
  -u, --url string      The target URL to submit the request to.

```

* output

This will just output the default smuggle gadgets. Currently the build in comfirmation only for H2.CL and H2.TE. But it is still possible to detect H2.0 attacks via timeouts (–filter TIMEOUT) or OAST on a domain of yours.

# Thanks (Without these people and projects this would not exist)

* [James Kettle](https://twitter.com/albinowax)
* [Emile Lerner](https://twitter.com/emil_lerner)
* [golang.org/x/net/http2](https://golang.org/x/net/http2)
* [golang.org/x/net/http2/hpack](https://golang.org/x/net/http2/hpack)
* [github.com/spf13/cobra](https://golang.org/x/net/http2/hpack)
