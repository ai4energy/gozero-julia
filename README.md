# gozero-julia
一个基于goctl的julia代码生成工具

开发中，内核还没写呢。

编译与运行
```go
go build -o gozero-julia.exe gozero-julia.go
go run gozero-julia.go --api demo/demo.api --dir outputdir
gozero-julia.exe --api demo/demo.api --dir outputdir
```

编译成共享库
```go
go build -o libgozero_julia.so -buildmode=c-shared exportdll.go
```

Julia调用
```julia
using Libdl
using JSON3
const lib = joinpath(@__DIR__, "libgozero_julia.so")

function JuliaGenCode(output_dir::String, api_file::String)::String
    res_ptr = ccall((:JuliaGenCode, lib), Cstring, (Cstring, Cstring), output_dir, api_file)
    return unsafe_string(res_ptr)
end

function GoZeroApiToJSON(api_file::String)::String
    res_ptr = ccall((:ExportToJSON, lib), Cstring, (Cstring,), api_file)
    result = unsafe_string(res_ptr)
    return result
end
output_dir = joinpath(@__DIR__, "ggoooodd")
api_file = joinpath(@__DIR__, "demo/demo.api")
result = JuliaGenCode(output_dir, api_file)
println(result)
json_str = GoZeroApiToJSON(api_file)

parsed_data = JSON3.read(json_str)
println(parsed_data)
```