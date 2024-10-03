function findIP(str) {

    let result = []
    const regex = /(\d+)\.(\d+)\.(\d+)\.(\d+)(:\d+)?/g

    const matches = [...str.matchAll(regex)];
    for (let match of matches) {
        let leadingZeros = false

        for (let j = 1; j <= 4; j++) {
            if (match[j].slice(0, 1) === '0' && match[j] != 0) {
                leadingZeros = true
            }
            if (match[j] == 0 && match[j].length>1) {
                leadingZeros = true
            }
        }
        if (!leadingZeros && match[1] < 256 && match[2] < 256 && match[3] < 256 && match[4] < 256) {
            if (match[5]) {
                if (match[5].slice(1) <= 65535) {
                    result.push(match[0])
                }
            } else {
                result.push(match[0])
            }
        }
    }
    console.log(result);
}

const dataSet = `qqq http:// qqqq q qqqqq https://something.com/hello qqqqqqq qhttp://example.com/hello?you=something&something=you qq 233.123.12.234 qw w wq wqw  wqw  ijnjjnfapsdbjnkfsdiqw klfsdjn fs fsd https://devdocs.io/javascript/global_objects/object/fromentries njnkfsdjnk sfdjn fsp fd192.168.1.123:8080 https://devdocs.io/javascript/global_objects/regexp/@@split
htpp://wrong/url hello %$& wf* ][½¬ http://correct/url?correct=yes è[}£§ https://nan-academy.github.io/js-training/?page=editor#data.nested 255.256.1233.2
ssages has become an accepted part of many cultures, as happened earlier with emailing. htt://[1] This makes texting a quick and http://www.example.com/mypage.html?crcat=test&crsource=test&crkw=buy-a-loteasy way to communicate 255.256.2 with friends, family and colleagues, including 255.256.555.2 in contexts where a call would be when one knows the other person is busy 192.169.1.23 with family or work activities).; 172.01.123.254:1234
for example, to order products or 10.1.23.7 http://www_example.com/ 255.255.255.000 09.09.09.09
services fromhttps://regex-performance.github.io/exercises.html 3...3 0.0.0.0:22 0.0.0.0:68768
this permits communication even between busy individuals255.253.123.2:8000 https: // . Text messages can also http:// be used to http://example.com/path?name=Branch&products=[Journeys,Email,Universal%20Ads]interact with automated systems,https:// regex -performance.github.io/ exercises.html172.01.123.999:1234
https//nan-academy.github.io/js-training/?page=editor#data.nested impolite or inappropriate (e.g., calling very late at night orhttp://localhost/exercises
https://192.168.1.123?something=nothing&pro=[23] htts:/nan-academy.github.io/js-training?b=123&a=123/?page=editor#data.nested  Like e-mail and voicemail and unlike calls (in which the caller hopes to speak directly with the recipient),
http://www.example.com/catalog.asp?itemid=232&template=fresh&crcat=ppc&crsource=google&crkw=buy-a-lot texting does not require the caller and recipient to both be free at the same moment0.0.0.0`

findIP(dataSet)