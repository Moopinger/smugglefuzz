
[![made-with-Go](https://img.shields.io/badge/made%20with-Go-brightgreen.svg)](http://golang.org)
[![go-report](https://goreportcard.com/badge/github.com/moopinger/smugglefuzz)](https://goreportcard.com/report/github.com/moopinger/smugglefuzz)
[![license](https://img.shields.io/badge/license-MIT-_red.svg)](https://opensource.org/licenses/MIT)


# SmuggleFuzz

SmuggleFuzz, a configurable and rapid HTTP downgrade smuggling scanner. Empowering users with customizable gadget lists, it offers deeper insights into the reasons behind failed smuggling attacks. I am confident that numerous HTTP downgrade-based smuggling vulnerabilities remain undiscovered. Unfortunately, many individuals cease their exploration after a cursory scan. By using SmuggleFuzz, you can uncover smuggling vulnerabilities that others may overlook.

For more info see [https://moopinger.github.io/blog/smugglefuzz/fuzzing/smuggling/2024/01/31/SmuggleFuzz.html](https://moopinger.github.io/blog/smugglefuzz/fuzzing/smuggling/2024/01/31/SmuggleFuzz.html)


# Install

Download a copy from the Releases page: [https://github.com/Moopinger/smugglefuzz/releases](https://github.com/Moopinger/smugglefuzz/releases)

# Or (Build yourself):

```
go install github.com/moopinger/smugglefuzz@latest
```


Another:


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
	smugglefuzz scan -u https://example.com/ --filter TIMEOUT --confirm
	smugglefuzz scan -u https://example.com/ -w wordlist.txt -t 10 --confirm --filter TIMEOUT
	smugglefuzz scan --dc -u https://example.com/ -w wordlist.txt -x PUT --confirm

	//Multiple targets? just use -f instead of -u and provide a file with the targets in it:

	smugglefuzz scan -f multiple_targets.txt --confirm -t 10 
	smugglefuzz scan -f multiple_targets.txt -w wordlist.txt --confirm -s ./save-success.txt --filter TIMEOUT
	smugglefuzz scan -f multiple_targets.txt -w wordlist.txt -H "Cookie: date=...; session=...;" -s ./save-success.txt -x PUT --filter TIMEOUT

Flags:
  -c, --confirm               Enable this flag to send a confirmation to the target when a timeout is encountered. Helps confirm if the target is vulnerable.
  -d, --data string           HTTP/2 Data frame to send. eg: 99\r\n (default "99\r\n")
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

Usage:
  smugglefuzz request [flags]

Examples:
	smugglefuzz request -u https://www.example.com/ -a "content-length\t; 13"
	smugglefuzz request -u https://www.example.com/ -a "content-length\t; 13" -x PUT
	smugglefuzz request -u https://www.example.com/ -a "content-length\t; 13" -H "Cookie: date=...; session=...;" --data "223\r\n"

Flags:
  -a, --attack string   Attack Header, separated by (; ) like the wordlist in 'scan' mode.
  -d, --data string     HTTP/2 Data frame content to send. (default "99\r\n")
      --dc              Disable colour in the output.
  -H, --header string   Insert custom header. eg "Cookie: values"
  -h, --help            help for request
  -i, --interval int    Detection timeout interval in seconds. (default 5)
  -x, --method string   The method to use. (default "POST")
  -u, --url string      The target URL to submit the request to.

```

* output

This will just output the default smuggle gadgets. Currently the built-in confirmation only works for H2.CL and H2.TE. But it is still possible to detect H2.0 attacks via timeouts (–filter TIMEOUT) or OAST on a domain of yours. (See the pseudo headers section below)


# Wordlist

SmuggleFuzz allows users complete control over requests through custom wordlists. These lists have a basic structure, which should be followed for optimal request handling. For instance, headers and their values are split using a `semicolon and a space (; )` instead of the usual colon, facilitating the inclusion of colon values in smuggling requests. This also opens up possibilities for various mutations and creative approaches. For detailed guidance on creating your own payloads, refer to: [James Kettle's HTTP/2 Research](https://portswigger.net/research/http2).

The tool includes a ready-to-use list of 125 smuggling gadgets, though there's always scope for expansion. These gadgets can be displayed using the 'output' command, providing insights into query structuring. Users can run scans with custom wordlists using the 'w' flag. The list supports URL encoding (%00) for non-printable byte values, such as carriage return and line feed represented as '%0d%0a' or '\r\n'. While the provided list is comprehensive, crafting your own gadgets can significantly enhance success rates.

# H2C

Need h2c support? Check out the "experimental-h2c" branch of SmuggleFuzz.

### Pseudo headers


Pseudo headers are fully supported and can be customized using the ":name" syntax:

* `:authority`
* `:scheme`
* `:method`

These headers are particularly useful when testing for H2.0 attack vectors. For example, in testing H2.0 smuggling using OAST-based methods, wordlist entries like the following can be effective:

As of v0.1.8 [HOSTNAME] can be used as a placeholder. SmuggleFuzz will replace this with the hostname at runtime.

* `:authority; [HOSTNAME]\r\n\r\nGET / HTTP/1.1\r\nHost: uniqid.oastify.com\r\nX-HEADER:%20`

or

* `:authority; [HOSTNAME]\r\n\r\nGET https://uniqid.oastify.com/ HTTP/1.1\r\nHost: [HOSTNAME]\r\nX-HEADER:%20`

# Detection

The detection method employed in SmuggleFuzz is akin to those used by Albinowax in his HTTP/1 smuggling research and defparam's Smuggler, albeit with adaptations for HTTP/2. Fortunately, as these attacks involve downgrades, the same request body/DATA payload can be utilized for both H2.CL and H2.TE detection.

A crucial guideline for wordlists is that successfully smuggled headers should include the following values for a confirmatory attack:

    CL: 13
    TE: chunked

Each request incorporates a "Body"/DATA-Frame with the value "99\r\n". If a TE header is smuggled, it is interpreted as a chunked content length of 153 (0x99) by the server if successfully smuggled, leading to a timeout. Similarly, smuggling a CL header with a content length of 13 also results in a timeout due to unfulfilled length.

The confirmatory request resubmits the same query that led to the timeout, but with Data: "3\r\nABC\r\n0\r\n\r\n". This satisfies both a chunked TE and a CL of 13, eliciting a successful response and potentially indicating a smuggling vulnerability.


# Thanks (Without these people and projects this would not exist)

* [James Kettle](https://twitter.com/albinowax)
* [Emile Lerner](https://twitter.com/emil_lerner)
* [golang.org/x/net/http2](https://golang.org/x/net/http2)
* [golang.org/x/net/http2/hpack](https://golang.org/x/net/http2/hpack)
* [github.com/spf13/cobra](https://golang.org/x/net/http2/hpack)
