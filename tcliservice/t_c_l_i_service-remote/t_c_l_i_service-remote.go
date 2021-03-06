// Autogenerated by Thrift Compiler (0.10.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
        "flag"
        "fmt"
        "math"
        "net"
        "net/url"
        "os"
        "strconv"
        "strings"
        "git.apache.org/thrift.git/lib/go/thrift"
        "github.com/weber09/gohive/tcliservice"
)


func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  TOpenSessionResp OpenSession(TOpenSessionReq req)")
  fmt.Fprintln(os.Stderr, "  TCloseSessionResp CloseSession(TCloseSessionReq req)")
  fmt.Fprintln(os.Stderr, "  TGetInfoResp GetInfo(TGetInfoReq req)")
  fmt.Fprintln(os.Stderr, "  TExecuteStatementResp ExecuteStatement(TExecuteStatementReq req)")
  fmt.Fprintln(os.Stderr, "  TGetTypeInfoResp GetTypeInfo(TGetTypeInfoReq req)")
  fmt.Fprintln(os.Stderr, "  TGetCatalogsResp GetCatalogs(TGetCatalogsReq req)")
  fmt.Fprintln(os.Stderr, "  TGetSchemasResp GetSchemas(TGetSchemasReq req)")
  fmt.Fprintln(os.Stderr, "  TGetTablesResp GetTables(TGetTablesReq req)")
  fmt.Fprintln(os.Stderr, "  TGetTableTypesResp GetTableTypes(TGetTableTypesReq req)")
  fmt.Fprintln(os.Stderr, "  TGetColumnsResp GetColumns(TGetColumnsReq req)")
  fmt.Fprintln(os.Stderr, "  TGetFunctionsResp GetFunctions(TGetFunctionsReq req)")
  fmt.Fprintln(os.Stderr, "  TGetPrimaryKeysResp GetPrimaryKeys(TGetPrimaryKeysReq req)")
  fmt.Fprintln(os.Stderr, "  TGetCrossReferenceResp GetCrossReference(TGetCrossReferenceReq req)")
  fmt.Fprintln(os.Stderr, "  TGetOperationStatusResp GetOperationStatus(TGetOperationStatusReq req)")
  fmt.Fprintln(os.Stderr, "  TCancelOperationResp CancelOperation(TCancelOperationReq req)")
  fmt.Fprintln(os.Stderr, "  TCloseOperationResp CloseOperation(TCloseOperationReq req)")
  fmt.Fprintln(os.Stderr, "  TGetResultSetMetadataResp GetResultSetMetadata(TGetResultSetMetadataReq req)")
  fmt.Fprintln(os.Stderr, "  TFetchResultsResp FetchResults(TFetchResultsReq req)")
  fmt.Fprintln(os.Stderr, "  TGetDelegationTokenResp GetDelegationToken(TGetDelegationTokenReq req)")
  fmt.Fprintln(os.Stderr, "  TCancelDelegationTokenResp CancelDelegationToken(TCancelDelegationTokenReq req)")
  fmt.Fprintln(os.Stderr, "  TRenewDelegationTokenResp RenewDelegationToken(TRenewDelegationTokenReq req)")
  fmt.Fprintln(os.Stderr)
  os.Exit(0)
}

func main() {
  flag.Usage = Usage
  var host string
  var port int
  var protocol string
  var urlString string
  var framed bool
  var useHttp bool
  var parsedUrl url.URL
  var trans thrift.TTransport
  _ = strconv.Atoi
  _ = math.Abs
  flag.Usage = Usage
  flag.StringVar(&host, "h", "localhost", "Specify host and port")
  flag.IntVar(&port, "p", 9090, "Specify port")
  flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
  flag.StringVar(&urlString, "u", "", "Specify the url")
  flag.BoolVar(&framed, "framed", false, "Use framed transport")
  flag.BoolVar(&useHttp, "http", false, "Use http")
  flag.Parse()
  
  if len(urlString) > 0 {
    parsedUrl, err := url.Parse(urlString)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
    host = parsedUrl.Host
    useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
  } else if useHttp {
    _, err := url.Parse(fmt.Sprint("http://", host, ":", port))
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
  }
  
  cmd := flag.Arg(0)
  var err error
  if useHttp {
    trans, err = thrift.NewTHttpClient(parsedUrl.String())
  } else {
    portStr := fmt.Sprint(port)
    if strings.Contains(host, ":") {
           host, portStr, err = net.SplitHostPort(host)
           if err != nil {
                   fmt.Fprintln(os.Stderr, "error with host:", err)
                   os.Exit(1)
           }
    }
    trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
    if err != nil {
      fmt.Fprintln(os.Stderr, "error resolving address:", err)
      os.Exit(1)
    }
    if framed {
      trans = thrift.NewTFramedTransport(trans)
    }
  }
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error creating transport", err)
    os.Exit(1)
  }
  defer trans.Close()
  var protocolFactory thrift.TProtocolFactory
  switch protocol {
  case "compact":
    protocolFactory = thrift.NewTCompactProtocolFactory()
    break
  case "simplejson":
    protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
    break
  case "json":
    protocolFactory = thrift.NewTJSONProtocolFactory()
    break
  case "binary", "":
    protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
    Usage()
    os.Exit(1)
  }
  client := tcliservice.NewTCLIServiceClientFactory(trans, protocolFactory)
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "OpenSession":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "OpenSession requires 1 args")
      flag.Usage()
    }
    arg74 := flag.Arg(1)
    mbTrans75 := thrift.NewTMemoryBufferLen(len(arg74))
    defer mbTrans75.Close()
    _, err76 := mbTrans75.WriteString(arg74)
    if err76 != nil {
      Usage()
      return
    }
    factory77 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt78 := factory77.GetProtocol(mbTrans75)
    argvalue0 := tcliservice.NewTOpenSessionReq()
    err79 := argvalue0.Read(jsProt78)
    if err79 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.OpenSession(value0))
    fmt.Print("\n")
    break
  case "CloseSession":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "CloseSession requires 1 args")
      flag.Usage()
    }
    arg80 := flag.Arg(1)
    mbTrans81 := thrift.NewTMemoryBufferLen(len(arg80))
    defer mbTrans81.Close()
    _, err82 := mbTrans81.WriteString(arg80)
    if err82 != nil {
      Usage()
      return
    }
    factory83 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt84 := factory83.GetProtocol(mbTrans81)
    argvalue0 := tcliservice.NewTCloseSessionReq()
    err85 := argvalue0.Read(jsProt84)
    if err85 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.CloseSession(value0))
    fmt.Print("\n")
    break
  case "GetInfo":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetInfo requires 1 args")
      flag.Usage()
    }
    arg86 := flag.Arg(1)
    mbTrans87 := thrift.NewTMemoryBufferLen(len(arg86))
    defer mbTrans87.Close()
    _, err88 := mbTrans87.WriteString(arg86)
    if err88 != nil {
      Usage()
      return
    }
    factory89 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt90 := factory89.GetProtocol(mbTrans87)
    argvalue0 := tcliservice.NewTGetInfoReq()
    err91 := argvalue0.Read(jsProt90)
    if err91 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetInfo(value0))
    fmt.Print("\n")
    break
  case "ExecuteStatement":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "ExecuteStatement requires 1 args")
      flag.Usage()
    }
    arg92 := flag.Arg(1)
    mbTrans93 := thrift.NewTMemoryBufferLen(len(arg92))
    defer mbTrans93.Close()
    _, err94 := mbTrans93.WriteString(arg92)
    if err94 != nil {
      Usage()
      return
    }
    factory95 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt96 := factory95.GetProtocol(mbTrans93)
    argvalue0 := tcliservice.NewTExecuteStatementReq()
    err97 := argvalue0.Read(jsProt96)
    if err97 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.ExecuteStatement(value0))
    fmt.Print("\n")
    break
  case "GetTypeInfo":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetTypeInfo requires 1 args")
      flag.Usage()
    }
    arg98 := flag.Arg(1)
    mbTrans99 := thrift.NewTMemoryBufferLen(len(arg98))
    defer mbTrans99.Close()
    _, err100 := mbTrans99.WriteString(arg98)
    if err100 != nil {
      Usage()
      return
    }
    factory101 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt102 := factory101.GetProtocol(mbTrans99)
    argvalue0 := tcliservice.NewTGetTypeInfoReq()
    err103 := argvalue0.Read(jsProt102)
    if err103 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetTypeInfo(value0))
    fmt.Print("\n")
    break
  case "GetCatalogs":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetCatalogs requires 1 args")
      flag.Usage()
    }
    arg104 := flag.Arg(1)
    mbTrans105 := thrift.NewTMemoryBufferLen(len(arg104))
    defer mbTrans105.Close()
    _, err106 := mbTrans105.WriteString(arg104)
    if err106 != nil {
      Usage()
      return
    }
    factory107 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt108 := factory107.GetProtocol(mbTrans105)
    argvalue0 := tcliservice.NewTGetCatalogsReq()
    err109 := argvalue0.Read(jsProt108)
    if err109 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetCatalogs(value0))
    fmt.Print("\n")
    break
  case "GetSchemas":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetSchemas requires 1 args")
      flag.Usage()
    }
    arg110 := flag.Arg(1)
    mbTrans111 := thrift.NewTMemoryBufferLen(len(arg110))
    defer mbTrans111.Close()
    _, err112 := mbTrans111.WriteString(arg110)
    if err112 != nil {
      Usage()
      return
    }
    factory113 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt114 := factory113.GetProtocol(mbTrans111)
    argvalue0 := tcliservice.NewTGetSchemasReq()
    err115 := argvalue0.Read(jsProt114)
    if err115 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetSchemas(value0))
    fmt.Print("\n")
    break
  case "GetTables":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetTables requires 1 args")
      flag.Usage()
    }
    arg116 := flag.Arg(1)
    mbTrans117 := thrift.NewTMemoryBufferLen(len(arg116))
    defer mbTrans117.Close()
    _, err118 := mbTrans117.WriteString(arg116)
    if err118 != nil {
      Usage()
      return
    }
    factory119 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt120 := factory119.GetProtocol(mbTrans117)
    argvalue0 := tcliservice.NewTGetTablesReq()
    err121 := argvalue0.Read(jsProt120)
    if err121 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetTables(value0))
    fmt.Print("\n")
    break
  case "GetTableTypes":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetTableTypes requires 1 args")
      flag.Usage()
    }
    arg122 := flag.Arg(1)
    mbTrans123 := thrift.NewTMemoryBufferLen(len(arg122))
    defer mbTrans123.Close()
    _, err124 := mbTrans123.WriteString(arg122)
    if err124 != nil {
      Usage()
      return
    }
    factory125 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt126 := factory125.GetProtocol(mbTrans123)
    argvalue0 := tcliservice.NewTGetTableTypesReq()
    err127 := argvalue0.Read(jsProt126)
    if err127 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetTableTypes(value0))
    fmt.Print("\n")
    break
  case "GetColumns":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetColumns requires 1 args")
      flag.Usage()
    }
    arg128 := flag.Arg(1)
    mbTrans129 := thrift.NewTMemoryBufferLen(len(arg128))
    defer mbTrans129.Close()
    _, err130 := mbTrans129.WriteString(arg128)
    if err130 != nil {
      Usage()
      return
    }
    factory131 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt132 := factory131.GetProtocol(mbTrans129)
    argvalue0 := tcliservice.NewTGetColumnsReq()
    err133 := argvalue0.Read(jsProt132)
    if err133 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetColumns(value0))
    fmt.Print("\n")
    break
  case "GetFunctions":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetFunctions requires 1 args")
      flag.Usage()
    }
    arg134 := flag.Arg(1)
    mbTrans135 := thrift.NewTMemoryBufferLen(len(arg134))
    defer mbTrans135.Close()
    _, err136 := mbTrans135.WriteString(arg134)
    if err136 != nil {
      Usage()
      return
    }
    factory137 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt138 := factory137.GetProtocol(mbTrans135)
    argvalue0 := tcliservice.NewTGetFunctionsReq()
    err139 := argvalue0.Read(jsProt138)
    if err139 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetFunctions(value0))
    fmt.Print("\n")
    break
  case "GetPrimaryKeys":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetPrimaryKeys requires 1 args")
      flag.Usage()
    }
    arg140 := flag.Arg(1)
    mbTrans141 := thrift.NewTMemoryBufferLen(len(arg140))
    defer mbTrans141.Close()
    _, err142 := mbTrans141.WriteString(arg140)
    if err142 != nil {
      Usage()
      return
    }
    factory143 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt144 := factory143.GetProtocol(mbTrans141)
    argvalue0 := tcliservice.NewTGetPrimaryKeysReq()
    err145 := argvalue0.Read(jsProt144)
    if err145 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetPrimaryKeys(value0))
    fmt.Print("\n")
    break
  case "GetCrossReference":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetCrossReference requires 1 args")
      flag.Usage()
    }
    arg146 := flag.Arg(1)
    mbTrans147 := thrift.NewTMemoryBufferLen(len(arg146))
    defer mbTrans147.Close()
    _, err148 := mbTrans147.WriteString(arg146)
    if err148 != nil {
      Usage()
      return
    }
    factory149 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt150 := factory149.GetProtocol(mbTrans147)
    argvalue0 := tcliservice.NewTGetCrossReferenceReq()
    err151 := argvalue0.Read(jsProt150)
    if err151 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetCrossReference(value0))
    fmt.Print("\n")
    break
  case "GetOperationStatus":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetOperationStatus requires 1 args")
      flag.Usage()
    }
    arg152 := flag.Arg(1)
    mbTrans153 := thrift.NewTMemoryBufferLen(len(arg152))
    defer mbTrans153.Close()
    _, err154 := mbTrans153.WriteString(arg152)
    if err154 != nil {
      Usage()
      return
    }
    factory155 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt156 := factory155.GetProtocol(mbTrans153)
    argvalue0 := tcliservice.NewTGetOperationStatusReq()
    err157 := argvalue0.Read(jsProt156)
    if err157 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetOperationStatus(value0))
    fmt.Print("\n")
    break
  case "CancelOperation":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "CancelOperation requires 1 args")
      flag.Usage()
    }
    arg158 := flag.Arg(1)
    mbTrans159 := thrift.NewTMemoryBufferLen(len(arg158))
    defer mbTrans159.Close()
    _, err160 := mbTrans159.WriteString(arg158)
    if err160 != nil {
      Usage()
      return
    }
    factory161 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt162 := factory161.GetProtocol(mbTrans159)
    argvalue0 := tcliservice.NewTCancelOperationReq()
    err163 := argvalue0.Read(jsProt162)
    if err163 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.CancelOperation(value0))
    fmt.Print("\n")
    break
  case "CloseOperation":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "CloseOperation requires 1 args")
      flag.Usage()
    }
    arg164 := flag.Arg(1)
    mbTrans165 := thrift.NewTMemoryBufferLen(len(arg164))
    defer mbTrans165.Close()
    _, err166 := mbTrans165.WriteString(arg164)
    if err166 != nil {
      Usage()
      return
    }
    factory167 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt168 := factory167.GetProtocol(mbTrans165)
    argvalue0 := tcliservice.NewTCloseOperationReq()
    err169 := argvalue0.Read(jsProt168)
    if err169 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.CloseOperation(value0))
    fmt.Print("\n")
    break
  case "GetResultSetMetadata":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetResultSetMetadata requires 1 args")
      flag.Usage()
    }
    arg170 := flag.Arg(1)
    mbTrans171 := thrift.NewTMemoryBufferLen(len(arg170))
    defer mbTrans171.Close()
    _, err172 := mbTrans171.WriteString(arg170)
    if err172 != nil {
      Usage()
      return
    }
    factory173 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt174 := factory173.GetProtocol(mbTrans171)
    argvalue0 := tcliservice.NewTGetResultSetMetadataReq()
    err175 := argvalue0.Read(jsProt174)
    if err175 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetResultSetMetadata(value0))
    fmt.Print("\n")
    break
  case "FetchResults":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "FetchResults requires 1 args")
      flag.Usage()
    }
    arg176 := flag.Arg(1)
    mbTrans177 := thrift.NewTMemoryBufferLen(len(arg176))
    defer mbTrans177.Close()
    _, err178 := mbTrans177.WriteString(arg176)
    if err178 != nil {
      Usage()
      return
    }
    factory179 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt180 := factory179.GetProtocol(mbTrans177)
    argvalue0 := tcliservice.NewTFetchResultsReq()
    err181 := argvalue0.Read(jsProt180)
    if err181 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.FetchResults(value0))
    fmt.Print("\n")
    break
  case "GetDelegationToken":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetDelegationToken requires 1 args")
      flag.Usage()
    }
    arg182 := flag.Arg(1)
    mbTrans183 := thrift.NewTMemoryBufferLen(len(arg182))
    defer mbTrans183.Close()
    _, err184 := mbTrans183.WriteString(arg182)
    if err184 != nil {
      Usage()
      return
    }
    factory185 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt186 := factory185.GetProtocol(mbTrans183)
    argvalue0 := tcliservice.NewTGetDelegationTokenReq()
    err187 := argvalue0.Read(jsProt186)
    if err187 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetDelegationToken(value0))
    fmt.Print("\n")
    break
  case "CancelDelegationToken":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "CancelDelegationToken requires 1 args")
      flag.Usage()
    }
    arg188 := flag.Arg(1)
    mbTrans189 := thrift.NewTMemoryBufferLen(len(arg188))
    defer mbTrans189.Close()
    _, err190 := mbTrans189.WriteString(arg188)
    if err190 != nil {
      Usage()
      return
    }
    factory191 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt192 := factory191.GetProtocol(mbTrans189)
    argvalue0 := tcliservice.NewTCancelDelegationTokenReq()
    err193 := argvalue0.Read(jsProt192)
    if err193 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.CancelDelegationToken(value0))
    fmt.Print("\n")
    break
  case "RenewDelegationToken":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "RenewDelegationToken requires 1 args")
      flag.Usage()
    }
    arg194 := flag.Arg(1)
    mbTrans195 := thrift.NewTMemoryBufferLen(len(arg194))
    defer mbTrans195.Close()
    _, err196 := mbTrans195.WriteString(arg194)
    if err196 != nil {
      Usage()
      return
    }
    factory197 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt198 := factory197.GetProtocol(mbTrans195)
    argvalue0 := tcliservice.NewTRenewDelegationTokenReq()
    err199 := argvalue0.Read(jsProt198)
    if err199 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.RenewDelegationToken(value0))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}
