# utoxss
A automede tool for mass cross site scripting(xss) hunting


<h4>utoxss</h4>
as you guys know if you went to use freq for xss you need qsreplace for adding payload on every endpoint then you can use freq but utoxss is again a advance version of freq + qsreplace so you just need to define your paylod and provide list it will automatice find cross site scripting for you.

<h5> usages:</h5>

``` waybackurls example.com | grep -v "?=" | utoxss -p '"><img src=x onerror=alert(1)>0' ```


<h5>installation</h5>

``` go install github.com/takshal/utoxss@latest ```
