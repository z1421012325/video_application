### 
短视频app项目
###


###
使用配置文件配置








### 



添加代码风格检测
https://github.com/golangci/golangci-lint-action



在yml文件中添加

    # 风格检查
    - name: golangci-lint
      uses: golangci/golangci-lint-action@v2
      with:
        version: v1.29


    # 安全检查
    - name: Run Gosec Security Scanner
      run:
        # 在自动编译的宿主机中下载gosec
        export PATH=$PATH:$(go env GOPATH)/bin
        go get github.com/securego/gosec
        gosec ./...
###