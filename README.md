Assignment
1. Build a simple URL shortener service that will accept a URL as an argument over a REST API and
return a shortened URL as a result.
a. If youhave not used orseena URL shortener before - please visit https://bitly.com/ and
try to shorten a URL. The goal of this assignment is not to build a fancy UI but an API
only version of that.
b. Don’t use a shortening API - you are supposed to write that part of code.
c. We expect your assignment to have a decent architecture.
2. The code should have following features:
a. If I again ask for the same URL,itshould give me the same URL as it gave before instead
of generating a new one.
b. If the user clicks on the shortURL then he should be redirected to the original URL. Write
a Redirection API that implements thisfunctionality.
c. The URL and shortened URL should be stored in-memory by application.
3. Write a metrics API that returns top 3 domain names that have been shortened the most
number of times. For eg. if the user hasshortened 4 YouTube video links, 1 StackOverflow link,2
Wikipedia links and 6 Udemy tutorial links. Then the output would be:
Udemy: 6
YouTube: 4
Wikipedia: 2
4. [BONUS] Putthis application in a Docker image by writing a Dockerfile and provide the docker
image link along with the source code link.



For Shortening----
Api - http://localhost:8084/api/shorten
REQ:
{
    "url":"https://www.google.com/hii"
}
RES:
{
    "success": true,
    "message": "Url shortened successfully",
    "shorturl": "XRlhcV"
}

For Redirecting----
Put the above response shorturl in the below api path
Api-http://localhost:8084/api/redirecturl/XRlhcV

For Metrics----
Api- http://localhost:8084/api/getmetrics
RES:
[
    {
        "domain": "www.google.com",
        "count": 9
    },
    {
        "domain": "github.com",
        "count": 1
    }
]





Docker Command: sudo docker build -t url-shortener . && sudo docker container run --restart unless-stopped -dit --name urlshortener --net=host --publish 8084:8084 url-shortener:latest
