
# Go 실행법
## vs code에서 실행
Ctrl+F5
	Debug > Run without Debugging 실행

## Go Command 실행
> go run  
```
$ go run [go 파일명]  
ex) go run hello.go
```

## Go Build 
> go build를 하면 실행 폴더에 exe 파일이 생성된다.
> 별도의 파일명이나 폴더명이 없어도 된다.
> 주의) go 파일명이 아닌 패키지명으로 exe 파일이 생성된다.
```
$ go build
```

## Go Build [파일명]
> 파일명으로 지정된 파일에 대해서만 exe 파일이 생성된다.
```
$ go build [파일명]  
ex) go build cal.go
```

## exe 파일 실행
```
[파일명.exe]로 바로 실행 가능
```