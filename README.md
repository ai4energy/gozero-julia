# gozero-julia
一个基于goctl的julia代码生成工具

开发中，内核还没写呢。

编译与运行
```go
go run gozero-julia.go julia --package somepackage --api demo/demo.api --dir outputdir
```

编译成共享库
```go
go build -o libgozero_julia.so -buildmode=c-shared exportdll.go
```

Julia调用
```julia
using Libdl
const lib = joinpath(@__DIR__, "libgozero_julia.so")

function JuliaGenCode(output_dir::String, api_file::String)::String
    res_ptr = ccall((:JuliaGenCode, lib), Cstring, (Cstring, Cstring), output_dir, api_file)
    return unsafe_string(res_ptr)
end

function GoZeroApiToJSON(output_dir::String, api_file::String)::String
    res_ptr = ccall((:ExportToJSON, lib), Cstring, (Cstring, Cstring), output_dir, api_file)
    result = unsafe_string(res_ptr)
    return result
end
output_dir = joinpath(@__DIR__, "ggoooodd")
api_file = joinpath(@__DIR__, "demo/demo.api")
result = JuliaGenCode(output_dir, api_file)
json_str = GoZeroApiToJSON(output_dir, api_file)
println(result)
parsed_data = JSON3.read(json_str)
```