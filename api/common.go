package gocast

const mainPage = `
<!DOCTYPE html>
<html>
    <head>
        <a href="/command/?Command=Play">Play</a>
        <a href="/command/?Command=Pause">Pause</a>
        <a href="/command/?Command=Next">Next</a>
    </head>
    <body>
        <ul>
            {{range .Files -}}
                <li> <a href="{{.Link}}">{{.Filename}}</a></li>
            {{- end -}}
        </ul>
    </body>
</html>
`

const switchSign = ";;"
const endSign = "\r\n\r\n"
