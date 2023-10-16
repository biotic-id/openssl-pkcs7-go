package interfaces

const not_qualified_ca = `
-----BEGIN CERTIFICATE-----
MIIHVTCCBwKgAwIBAgIQAdnGAyzbboAAAAAANcQAATAKBggqhQMHAQEDAjBgMQsw
CQYDVQQGEwJSVTFRME8GA1UEAwxI0KHQv9C10YbQuNCw0LvQuNC30LjRgNC+0LLQ
sNC90L3Ri9C5INGG0LXQvdGC0YAg0YHQtdGA0YLQuNGE0LjQutCw0YbQuNC4MB4X
DTIzMDgwMzEyMDgwMFoXDTM4MTExNjA4NTEwMFowgYQxCzAJBgNVBAYTAlJVMQow
CAYDVQQLDAExMWkwZwYDVQQDDGDQk9C+0YHRg9GB0LvRg9Cz0LguINCd0LXQutCy
0LDQu9C40YTQuNGG0LjRgNC+0LLQsNC90L3QsNGPINGN0LvQtdC60YLRgNC+0L3Q
vdCw0Y8g0L/QvtC00L/QuNGB0YwwZjAfBggqhQMHAQEBATATBgcqhQMCAiMBBggq
hQMHAQECAgNDAARAwhA/6i62zJKZoB2Rk6/fQ568JXH6N34KhnhiqhuUXefMNZPa
Cv8eyZuZi/LQCyKqwgGvMDgWQwpvo0lBPRa5HoEJADM1QzQwMDAxo4IFXzCCBVsw
KwYDVR0QBCQwIoAPMjAyMzA4MDMxMjA4MDBagQ8yMDI0MTEwMzEyMDgwMFowDgYD
VR0PAQH/BAQDAgHGMA8GA1UdEwEB/wQFMAMBAf8wHQYDVR0OBBYEFKaQBcnMV1Qd
S0F1X4o7/FiK2p7VMIGoBgUqhQNkbwSBngyBm9Ch0YDQtdC00YHRgtCy0L4g0LrR
gNC40L/RgtC+0LPRgNCw0YTQuNGH0LXRgdC60L7QuSDQt9Cw0YnQuNGC0Ysg0LjQ
vdGE0L7RgNC80LDRhtC40Lgg0KHQmtCX0JggVmlQTmV0IENTUCA0LjQgKNCS0LXR
gNGB0LjRjyA0LjQuMikgKNC40YHQv9C+0LvQvdC10L3QuNC1IDMpMIIBzQYFKoUD
ZHAEggHCMIIBvgyBm9Ch0YDQtdC00YHRgtCy0L4g0LrRgNC40L/RgtC+0LPRgNCw
0YTQuNGH0LXRgdC60L7QuSDQt9Cw0YnQuNGC0Ysg0LjQvdGE0L7RgNC80LDRhtC4
0Lgg0KHQmtCX0JggVmlQTmV0IENTUCA0LjQgKNCS0LXRgNGB0LjRjyA0LjQuMikg
KNC40YHQv9C+0LvQvdC10L3QuNC1IDMpDG3Qn9GA0L7Qs9GA0LDQvNC80L3Ri9C5
INC60L7QvNC/0LvQtdC60YEgIlZpUE5ldCDQo9C00L7RgdGC0L7QstC10YDRj9GO
0YnQuNC5INGG0LXQvdGC0YAgNCAo0LLQtdGA0YHQuNGPIDQuNikiDGTQodC10YDR
gtC40YTQuNC60LDRgiDRgdC+0L7RgtCy0LXRgtGB0YLQstC40Y8g4oSWINCh0KQv
MTI0LTQxMDMg0L7RgiAxMCDQsNCy0LPRg9GB0YLQsCAyMDIxINCz0L7QtNCwDEnQ
l9Cw0LrQu9GO0YfQtdC90LjQtSDihJYgMTQ5LzcvNi80NDYg0L7RgiAyOSDQtNC1
0LrQsNCx0YDRjyAyMDIxINCz0L7QtNCwMIHgBggrBgEFBQcBAQSB0zCB0DAzBggr
BgEFBQcwAoYnaHR0cDovL2hvc3QxLmdvc2tleS5ydS9DQS9TQ0MxLTIwMjIuY2Vy
MDEGCCsGAQUFBzAChiVodHRwOi8vMTkzLjE2Mi4zMC44Ni9DQS9TQ0MxLTIwMjIu
Y2VyMDMGCCsGAQUFBzAChidodHRwOi8vaG9zdDIuZ29za2V5LnJ1L0NBL1NDQzEt
MjAyMi5jZXIwMQYIKwYBBQUHMAKGJWh0dHA6Ly8xOTMuMTYyLjMwLjczL0NBL1ND
QzEtMjAyMi5jZXIwgccGA1UdHwSBvzCBvDAuoCygKoYoaHR0cDovL2hvc3QxLmdv
c2tleS5ydS9DUkwvU0NDMS0yMDIyLmNybDAsoCqgKIYmaHR0cDovLzE5My4xNjIu
MzAuODYvQ1JML1NDQzEtMjAyMi5jcmwwLqAsoCqGKGh0dHA6Ly9ob3N0Mi5nb3Nr
ZXkucnUvQ1JML1NDQzEtMjAyMi5jcmwwLKAqoCiGJmh0dHA6Ly8xOTMuMTYyLjMw
LjczL0NSTC9TQ0MxLTIwMjIuY3JsMIGZBgNVHSMEgZEwgY6AFIuKyk0LqD8YP8tV
t0/48mjlMvCKoWSkYjBgMQswCQYDVQQGEwJSVTFRME8GA1UEAwxI0KHQv9C10YbQ
uNCw0LvQuNC30LjRgNC+0LLQsNC90L3Ri9C5INGG0LXQvdGC0YAg0YHQtdGA0YLQ
uNGE0LjQutCw0YbQuNC4ghAB2PmYosgjAAAAAAA1xAABMCcGA1UdIAQgMB4wCAYG
KoUDZHEBMAgGBiqFA2RxAjAIBgYqhQNkcQMwCgYIKoUDBwEBAwIDQQBUS8VWrAdV
CwAcpeeTqdGRkEFAuGvynMK9YzcKS2QnYiphRSF1PrgWxVBUi6uADacGFPZtXJwE
cENPiG2yv11b
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIIFjjCCBTugAwIBAgIQAdj5mKLIIwAAAAAANcQAATAKBggqhQMHAQEDAjBgMQsw
CQYDVQQGEwJSVTFRME8GA1UEAwxI0KHQv9C10YbQuNCw0LvQuNC30LjRgNC+0LLQ
sNC90L3Ri9C5INGG0LXQvdGC0YAg0YHQtdGA0YLQuNGE0LjQutCw0YbQuNC4MB4X
DTIyMTExNjA4NTEwMFoXDTM4MTExNjA4NTEwMFowYDELMAkGA1UEBhMCUlUxUTBP
BgNVBAMMSNCh0L/QtdGG0LjQsNC70LjQt9C40YDQvtCy0LDQvdC90YvQuSDRhtC1
0L3RgtGAINGB0LXRgNGC0LjRhNC40LrQsNGG0LjQuDBmMB8GCCqFAwcBAQEBMBMG
ByqFAwICIwEGCCqFAwcBAQICA0MABEB1mKFsIch13QtROb/E/oEry1arkU/f0vNV
n/vrW8fXZMUP1e8lky57nFnQ0X1yoCDXa/nEdt/UMgUVFvHpJ0IJgQkAMzVDNDAw
MDGCCQAzNUM0MDAwMaOCA7IwggOuMA4GA1UdDwEB/wQEAwIBxjAPBgNVHRMBAf8E
BTADAQH/MIGoBgUqhQNkbwSBngyBm9Ch0YDQtdC00YHRgtCy0L4g0LrRgNC40L/R
gtC+0LPRgNCw0YTQuNGH0LXRgdC60L7QuSDQt9Cw0YnQuNGC0Ysg0LjQvdGE0L7R
gNC80LDRhtC40Lgg0KHQmtCX0JggVmlQTmV0IENTUCA0LjQgKNCS0LXRgNGB0LjR
jyA0LjQuMikgKNC40YHQv9C+0LvQvdC10L3QuNC1IDMpMIIBzQYFKoUDZHAEggHC
MIIBvgyBm9Ch0YDQtdC00YHRgtCy0L4g0LrRgNC40L/RgtC+0LPRgNCw0YTQuNGH
0LXRgdC60L7QuSDQt9Cw0YnQuNGC0Ysg0LjQvdGE0L7RgNC80LDRhtC40Lgg0KHQ
mtCX0JggVmlQTmV0IENTUCA0LjQgKNCS0LXRgNGB0LjRjyA0LjQuMikgKNC40YHQ
v9C+0LvQvdC10L3QuNC1IDMpDG3Qn9GA0L7Qs9GA0LDQvNC80L3Ri9C5INC60L7Q
vNC/0LvQtdC60YEgIlZpUE5ldCDQo9C00L7RgdGC0L7QstC10YDRj9GO0YnQuNC5
INGG0LXQvdGC0YAgNCAo0LLQtdGA0YHQuNGPIDQuNikiDGTQodC10YDRgtC40YTQ
uNC60LDRgiDRgdC+0L7RgtCy0LXRgtGB0YLQstC40Y8g4oSWINCh0KQvMTI0LTQx
MDMg0L7RgiAxMCDQsNCy0LPRg9GB0YLQsCAyMDIxINCz0L7QtNCwDEnQl9Cw0LrQ
u9GO0YfQtdC90LjQtSDihJYgMTQ5LzcvNi80NDYg0L7RgiAyOSDQtNC10LrQsNCx
0YDRjyAyMDIxINCz0L7QtNCwMCcGA1UdIAQgMB4wCAYGKoUDZHEBMAgGBiqFA2Rx
AjAIBgYqhQNkcQMwKwYDVR0QBCQwIoAPMjAyMjExMTYwODUxMDBagQ8yMDI4MTEx
NjA4NTEwMFowHQYDVR0OBBYEFIuKyk0LqD8YP8tVt0/48mjlMvCKMIGZBgNVHSME
gZEwgY6AFIuKyk0LqD8YP8tVt0/48mjlMvCKoWSkYjBgMQswCQYDVQQGEwJSVTFR
ME8GA1UEAwxI0KHQv9C10YbQuNCw0LvQuNC30LjRgNC+0LLQsNC90L3Ri9C5INGG
0LXQvdGC0YAg0YHQtdGA0YLQuNGE0LjQutCw0YbQuNC4ghAB2PmYosgjAAAAAAA1
xAABMAoGCCqFAwcBAQMCA0EA7a87iNw4Wx2oFDMrBGe8OFEEQUMsDay6P3iPnaoD
kzyDOaox5ChNe/w45erO7QFjCJ6u6j3r/dQrQltqwOEJzg==
-----END CERTIFICATE-----
`
const qualified_ca = `
-----BEGIN CERTIFICATE-----
MIIIGjCCB8egAwIBAgILAMb6ZbYAAAAABnMwCgYIKoUDBwEBAwIwggE7MSEwHwYJ
KoZIhvcNAQkBFhJkaXRAZGlnaXRhbC5nb3YucnUxCzAJBgNVBAYTAlJVMRgwFgYD
VQQIDA83NyDQnNC+0YHQutCy0LAxGTAXBgNVBAcMENCzLiDQnNC+0YHQutCy0LAx
UzBRBgNVBAkMStCf0YDQtdGB0L3QtdC90YHQutCw0Y8g0L3QsNCx0LXRgNC10LbQ
vdCw0Y8sINC00L7QvCAxMCwg0YHRgtGA0L7QtdC90LjQtSAyMSYwJAYDVQQKDB3Q
nNC40L3RhtC40YTRgNGLINCg0L7RgdGB0LjQuDEYMBYGBSqFA2QBEg0xMDQ3NzAy
MDI2NzAxMRUwEwYFKoUDZAQSCjc3MTA0NzQzNzUxJjAkBgNVBAMMHdCc0LjQvdGG
0LjRhNGA0Ysg0KDQvtGB0YHQuNC4MB4XDTIyMDYyOTEzMTEwOFoXDTM3MDYyOTEz
MTEwOFowggFkMSUwIwYJKoZIhvcNAQkBFhZTdXBwb3J0SUlUQGluZm90ZWNzLnJ1
MU8wTQYDVQQJDEbRg9C7LiDQnNC40YjQuNC90LAsINC0LiA1Niwg0YHRgtGALiAy
LCDRjdGCLiAyLCDQv9C+0LwuIElYLCDQutC+0LwuIDExMQswCQYDVQQGEwJSVTEc
MBoGA1UECAwTNzcg0LMuINCc0L7RgdC60LLQsDEVMBMGA1UEBwwM0JzQvtGB0LrQ
stCwMV8wXQYDVQQKDFbQkNC60YbQuNC+0L3QtdGA0L3QvtC1INCe0LHRidC10YHR
gtCy0L4gItCY0L3RhNC+0KLQtdCa0KEg0JjQvdGC0LXRgNC90LXRgiDQotGA0LDR
gdGCIjEYMBYGBSqFA2QBEg0xMDI3NzM5MTEzMDQ5MRUwEwYFKoUDZAQSCjc3NDMw
MjA1NjAxFjAUBgNVBAMMDdCQ0J4gItCY0JjQoiIwZjAfBggqhQMHAQEBATATBgcq
hQMCAiMBBggqhQMHAQECAgNDAARAPBcayeXjg9kg6+H/VIDJTWw3D111+MbrD0EA
muizMfB1UIkRtjhNt7mB+5gwbA4ZtTBBeIC4X3sd/DmaBXsyu6OCBHYwggRyMA4G
A1UdDwEB/wQEAwIBxjASBgNVHRMBAf8ECDAGAQH/AgEAMIGoBgUqhQNkbwSBngyB
m9Ch0YDQtdC00YHRgtCy0L4g0LrRgNC40L/RgtC+0LPRgNCw0YTQuNGH0LXRgdC6
0L7QuSDQt9Cw0YnQuNGC0Ysg0LjQvdGE0L7RgNC80LDRhtC40Lgg0KHQmtCX0Jgg
VmlQTmV0IENTUCA0LjQgKNCS0LXRgNGB0LjRjyA0LjQuMikgKNC40YHQv9C+0LvQ
vdC10L3QuNC1IDMpMCcGA1UdIAQgMB4wCAYGKoUDZHEBMAgGBiqFA2RxAjAIBgYq
hQNkcQMwHQYDVR0OBBYEFCz3/6Qvrn/uqibKHxU0NQ1MT4cpMIIBfQYDVR0jBIIB
dDCCAXCAFMkTWLFMp2I6ftI/PKbnFHydcKOGoYIBQ6SCAT8wggE7MSEwHwYJKoZI
hvcNAQkBFhJkaXRAZGlnaXRhbC5nb3YucnUxCzAJBgNVBAYTAlJVMRgwFgYDVQQI
DA83NyDQnNC+0YHQutCy0LAxGTAXBgNVBAcMENCzLiDQnNC+0YHQutCy0LAxUzBR
BgNVBAkMStCf0YDQtdGB0L3QtdC90YHQutCw0Y8g0L3QsNCx0LXRgNC10LbQvdCw
0Y8sINC00L7QvCAxMCwg0YHRgtGA0L7QtdC90LjQtSAyMSYwJAYDVQQKDB3QnNC4
0L3RhtC40YTRgNGLINCg0L7RgdGB0LjQuDEYMBYGBSqFA2QBEg0xMDQ3NzAyMDI2
NzAxMRUwEwYFKoUDZAQSCjc3MTA0NzQzNzUxJjAkBgNVBAMMHdCc0LjQvdGG0LjR
hNGA0Ysg0KDQvtGB0YHQuNC4ghEAlR+jR3xhBDqt+oWGJ4I0QjCBjwYDVR0fBIGH
MIGEMCqgKKAmhiRodHRwOi8vcmVlc3RyLXBraS5ydS9jZHAvZ3VjMjAyMi5jcmww
KqAooCaGJGh0dHA6Ly9jb21wYW55LnJ0LnJ1L2NkcC9ndWMyMDIyLmNybDAqoCig
JoYkaHR0cDovL3Jvc3RlbGVjb20ucnUvY2RwL2d1YzIwMjIuY3JsMEAGCCsGAQUF
BwEBBDQwMjAwBggrBgEFBQcwAoYkaHR0cDovL3JlZXN0ci1wa2kucnUvY2RwL2d1
YzIwMjIuY3J0MIH1BgUqhQNkcASB6zCB6Aw00J/QkNCa0JwgwqvQmtGA0LjQv9GC
0L7Qn9GA0L4gSFNNwrsg0LLQtdGA0YHQuNC4IDIuMAxD0J/QkNCaIMKr0JPQvtC7
0L7QstC90L7QuSDRg9C00L7RgdGC0L7QstC10YDRj9GO0YnQuNC5INGG0LXQvdGC
0YDCuww10JfQsNC60LvRjtGH0LXQvdC40LUg4oSWIDE0OS8zLzIvMi8yMyDQvtGC
IDAyLjAzLjIwMTgMNNCX0LDQutC70Y7Rh9C10L3QuNC1IOKEliAxNDkvNy82LTQ0
OSDQvtGCIDMwLjEyLjIwMjEwDAYFKoUDZHIEAwIBATAKBggqhQMHAQEDAgNBAAmA
7RtNvsqSE01+uO2fdDHe+VHMkHvMQaYxY3t+qT1ph8mP2fWSUJyZwPhjx3GBAPXq
i2RPfjB23oMqz3h0keo=
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIIFUTCCBP6gAwIBAgIRAJUfo0d8YQQ6rfqFhieCNEIwCgYIKoUDBwEBAwIwggE7
MSEwHwYJKoZIhvcNAQkBFhJkaXRAZGlnaXRhbC5nb3YucnUxCzAJBgNVBAYTAlJV
MRgwFgYDVQQIDA83NyDQnNC+0YHQutCy0LAxGTAXBgNVBAcMENCzLiDQnNC+0YHQ
utCy0LAxUzBRBgNVBAkMStCf0YDQtdGB0L3QtdC90YHQutCw0Y8g0L3QsNCx0LXR
gNC10LbQvdCw0Y8sINC00L7QvCAxMCwg0YHRgtGA0L7QtdC90LjQtSAyMSYwJAYD
VQQKDB3QnNC40L3RhtC40YTRgNGLINCg0L7RgdGB0LjQuDEYMBYGBSqFA2QBEg0x
MDQ3NzAyMDI2NzAxMRUwEwYFKoUDZAQSCjc3MTA0NzQzNzUxJjAkBgNVBAMMHdCc
0LjQvdGG0LjRhNGA0Ysg0KDQvtGB0YHQuNC4MB4XDTIyMDEwODEzMzIzOVoXDTQw
MDEwODEzMzIzOVowggE7MSEwHwYJKoZIhvcNAQkBFhJkaXRAZGlnaXRhbC5nb3Yu
cnUxCzAJBgNVBAYTAlJVMRgwFgYDVQQIDA83NyDQnNC+0YHQutCy0LAxGTAXBgNV
BAcMENCzLiDQnNC+0YHQutCy0LAxUzBRBgNVBAkMStCf0YDQtdGB0L3QtdC90YHQ
utCw0Y8g0L3QsNCx0LXRgNC10LbQvdCw0Y8sINC00L7QvCAxMCwg0YHRgtGA0L7Q
tdC90LjQtSAyMSYwJAYDVQQKDB3QnNC40L3RhtC40YTRgNGLINCg0L7RgdGB0LjQ
uDEYMBYGBSqFA2QBEg0xMDQ3NzAyMDI2NzAxMRUwEwYFKoUDZAQSCjc3MTA0NzQz
NzUxJjAkBgNVBAMMHdCc0LjQvdGG0LjRhNGA0Ysg0KDQvtGB0YHQuNC4MGYwHwYI
KoUDBwEBAQEwEwYHKoUDAgIjAQYIKoUDBwEBAgIDQwAEQFpKa6Qda48LjFq/drz2
M27fFqu/g3+Prxrg9lE+KPzdbvRHwuOtbhlJ92ogmS+i7mhDDGPsWhtPVduV9Kbq
QI6jggHQMIIBzDCB9QYFKoUDZHAEgeswgegMNNCf0JDQmtCcIMKr0JrRgNC40L/R
gtC+0J/RgNC+IEhTTcK7INCy0LXRgNGB0LjQuCAyLjAMQ9Cf0JDQmiDCq9CT0L7Q
u9C+0LLQvdC+0Lkg0YPQtNC+0YHRgtC+0LLQtdGA0Y/RjtGJ0LjQuSDRhtC10L3R
gtGAwrsMNdCX0LDQutC70Y7Rh9C10L3QuNC1IOKEliAxNDkvMy8yLzIvMjMg0L7R
giAwMi4wMy4yMDE4DDTQl9Cw0LrQu9GO0YfQtdC90LjQtSDihJYgMTQ5LzcvNi00
NDkg0L7RgiAzMC4xMi4yMDIxMD8GBSqFA2RvBDYMNNCf0JDQmtCcIMKr0JrRgNC4
0L/RgtC+0J/RgNC+IEhTTcK7INCy0LXRgNGB0LjQuCAyLjAwDAYFKoUDZHIEAwIB
ADBDBgNVHSAEPDA6MAgGBiqFA2RxATAIBgYqhQNkcQIwCAYGKoUDZHEDMAgGBiqF
A2RxBDAIBgYqhQNkcQUwBgYEVR0gADAOBgNVHQ8BAf8EBAMCAQYwDwYDVR0TAQH/
BAUwAwEB/zAdBgNVHQ4EFgQUyRNYsUynYjp+0j88pucUfJ1wo4YwCgYIKoUDBwEB
AwIDQQCCSXhICg3SZ/TTCtRJpBFuXGSy3PeZTEeOwdOHIv0tWiN2q0mPRzB/o6r9
MXjGqdzfYGtCrq1l5FsXZOI5c/2S
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIIHcjCCBx+gAwIBAgILAPjJxAQAAAAABfIwCgYIKoUDBwEBAwIwggEkMR4wHAYJ
KoZIhvcNAQkBFg9kaXRAbWluc3Z5YXoucnUxCzAJBgNVBAYTAlJVMRgwFgYDVQQI
DA83NyDQnNC+0YHQutCy0LAxGTAXBgNVBAcMENCzLiDQnNC+0YHQutCy0LAxLjAs
BgNVBAkMJdGD0LvQuNGG0LAg0KLQstC10YDRgdC60LDRjywg0LTQvtC8IDcxLDAq
BgNVBAoMI9Cc0LjQvdC60L7QvNGB0LLRj9C30Ywg0KDQvtGB0YHQuNC4MRgwFgYF
KoUDZAESDTEwNDc3MDIwMjY3MDExGjAYBggqhQMDgQMBARIMMDA3NzEwNDc0Mzc1
MSwwKgYDVQQDDCPQnNC40L3QutC+0LzRgdCy0Y/Qt9GMINCg0L7RgdGB0LjQuDAe
Fw0yMjAxMDgxNTQ0MjBaFw0zNzAxMDgxNTQ0MjBaMIIBOzEhMB8GCSqGSIb3DQEJ
ARYSZGl0QGRpZ2l0YWwuZ292LnJ1MQswCQYDVQQGEwJSVTEYMBYGA1UECAwPNzcg
0JzQvtGB0LrQstCwMRkwFwYDVQQHDBDQsy4g0JzQvtGB0LrQstCwMVMwUQYDVQQJ
DErQn9GA0LXRgdC90LXQvdGB0LrQsNGPINC90LDQsdC10YDQtdC20L3QsNGPLCDQ
tNC+0LwgMTAsINGB0YLRgNC+0LXQvdC40LUgMjEmMCQGA1UECgwd0JzQuNC90YbQ
uNGE0YDRiyDQoNC+0YHRgdC40LgxGDAWBgUqhQNkARINMTA0NzcwMjAyNjcwMTEV
MBMGBSqFA2QEEgo3NzEwNDc0Mzc1MSYwJAYDVQQDDB3QnNC40L3RhtC40YTRgNGL
INCg0L7RgdGB0LjQuDBmMB8GCCqFAwcBAQEBMBMGByqFAwICIwEGCCqFAwcBAQIC
A0MABEBaSmukHWuPC4xav3a89jNu3xarv4N/j68a4PZRPij83W70R8LjrW4ZSfdq
IJkvou5oQwxj7FobT1XblfSm6kCOo4IEDjCCBAowPwYFKoUDZG8ENgw00J/QkNCa
0JwgwqvQmtGA0LjQv9GC0L7Qn9GA0L4gSFNNwrsg0LLQtdGA0YHQuNC4IDIuMDAM
BgUqhQNkcgQDAgEAMEMGA1UdIAQ8MDowCAYGKoUDZHEBMAgGBiqFA2RxAjAIBgYq
hQNkcQMwCAYGKoUDZHEEMAgGBiqFA2RxBTAGBgRVHSAAMA4GA1UdDwEB/wQEAwIB
BjAPBgNVHRMBAf8EBTADAQH/MB0GA1UdDgQWBBTJE1ixTKdiOn7SPzym5xR8nXCj
hjCCAWYGA1UdIwSCAV0wggFZgBQZd4+7guZvyF+T8BUdkyKh1q0MJqGCASykggEo
MIIBJDEeMBwGCSqGSIb3DQEJARYPZGl0QG1pbnN2eWF6LnJ1MQswCQYDVQQGEwJS
VTEYMBYGA1UECAwPNzcg0JzQvtGB0LrQstCwMRkwFwYDVQQHDBDQsy4g0JzQvtGB
0LrQstCwMS4wLAYDVQQJDCXRg9C70LjRhtCwINCi0LLQtdGA0YHQutCw0Y8sINC0
0L7QvCA3MSwwKgYDVQQKDCPQnNC40L3QutC+0LzRgdCy0Y/Qt9GMINCg0L7RgdGB
0LjQuDEYMBYGBSqFA2QBEg0xMDQ3NzAyMDI2NzAxMRowGAYIKoUDA4EDAQESDDAw
NzcxMDQ3NDM3NTEsMCoGA1UEAwwj0JzQuNC90LrQvtC80YHQstGP0LfRjCDQoNC+
0YHRgdC40LiCEQDqyjLO9fl51o08Tk8sxoekMIGPBgNVHR8EgYcwgYQwKqAooCaG
JGh0dHA6Ly9yZWVzdHItcGtpLnJ1L2NkcC9ndWMyMDIxLmNybDAqoCigJoYkaHR0
cDovL2NvbXBhbnkucnQucnUvY2RwL2d1YzIwMjEuY3JsMCqgKKAmhiRodHRwOi8v
cm9zdGVsZWNvbS5ydS9jZHAvZ3VjMjAyMS5jcmwwQAYIKwYBBQUHAQEENDAyMDAG
CCsGAQUFBzAChiRodHRwOi8vcmVlc3RyLXBraS5ydS9jZHAvZ3VjMjAyMS5jcnQw
gfUGBSqFA2RwBIHrMIHoDDTQn9CQ0JrQnCDCq9Ca0YDQuNC/0YLQvtCf0YDQviBI
U03CuyDQstC10YDRgdC40LggMi4wDEPQn9CQ0JogwqvQk9C+0LvQvtCy0L3QvtC5
INGD0LTQvtGB0YLQvtCy0LXRgNGP0Y7RidC40Lkg0YbQtdC90YLRgMK7DDXQl9Cw
0LrQu9GO0YfQtdC90LjQtSDihJYgMTQ5LzMvMi8yLzIzINC+0YIgMDIuMDMuMjAx
OAw00JfQsNC60LvRjtGH0LXQvdC40LUg4oSWIDE0OS83LzYtNDQ5INC+0YIgMzAu
MTIuMjAyMTAKBggqhQMHAQEDAgNBAKUHyIN2tDVgKi1xtV4r3Y7kMx+Ycxkg9Ysb
yc3hAwffaML3S/SGaDopCQgKoyU8gJZ7LoHA4VHXFhwpID/oa7s=
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIIFUTCCBP6gAwIBAgIRAJUfo0d8YQQ6rfqFhieCNEIwCgYIKoUDBwEBAwIwggE7
MSEwHwYJKoZIhvcNAQkBFhJkaXRAZGlnaXRhbC5nb3YucnUxCzAJBgNVBAYTAlJV
MRgwFgYDVQQIDA83NyDQnNC+0YHQutCy0LAxGTAXBgNVBAcMENCzLiDQnNC+0YHQ
utCy0LAxUzBRBgNVBAkMStCf0YDQtdGB0L3QtdC90YHQutCw0Y8g0L3QsNCx0LXR
gNC10LbQvdCw0Y8sINC00L7QvCAxMCwg0YHRgtGA0L7QtdC90LjQtSAyMSYwJAYD
VQQKDB3QnNC40L3RhtC40YTRgNGLINCg0L7RgdGB0LjQuDEYMBYGBSqFA2QBEg0x
MDQ3NzAyMDI2NzAxMRUwEwYFKoUDZAQSCjc3MTA0NzQzNzUxJjAkBgNVBAMMHdCc
0LjQvdGG0LjRhNGA0Ysg0KDQvtGB0YHQuNC4MB4XDTIyMDEwODEzMzIzOVoXDTQw
MDEwODEzMzIzOVowggE7MSEwHwYJKoZIhvcNAQkBFhJkaXRAZGlnaXRhbC5nb3Yu
cnUxCzAJBgNVBAYTAlJVMRgwFgYDVQQIDA83NyDQnNC+0YHQutCy0LAxGTAXBgNV
BAcMENCzLiDQnNC+0YHQutCy0LAxUzBRBgNVBAkMStCf0YDQtdGB0L3QtdC90YHQ
utCw0Y8g0L3QsNCx0LXRgNC10LbQvdCw0Y8sINC00L7QvCAxMCwg0YHRgtGA0L7Q
tdC90LjQtSAyMSYwJAYDVQQKDB3QnNC40L3RhtC40YTRgNGLINCg0L7RgdGB0LjQ
uDEYMBYGBSqFA2QBEg0xMDQ3NzAyMDI2NzAxMRUwEwYFKoUDZAQSCjc3MTA0NzQz
NzUxJjAkBgNVBAMMHdCc0LjQvdGG0LjRhNGA0Ysg0KDQvtGB0YHQuNC4MGYwHwYI
KoUDBwEBAQEwEwYHKoUDAgIjAQYIKoUDBwEBAgIDQwAEQFpKa6Qda48LjFq/drz2
M27fFqu/g3+Prxrg9lE+KPzdbvRHwuOtbhlJ92ogmS+i7mhDDGPsWhtPVduV9Kbq
QI6jggHQMIIBzDCB9QYFKoUDZHAEgeswgegMNNCf0JDQmtCcIMKr0JrRgNC40L/R
gtC+0J/RgNC+IEhTTcK7INCy0LXRgNGB0LjQuCAyLjAMQ9Cf0JDQmiDCq9CT0L7Q
u9C+0LLQvdC+0Lkg0YPQtNC+0YHRgtC+0LLQtdGA0Y/RjtGJ0LjQuSDRhtC10L3R
gtGAwrsMNdCX0LDQutC70Y7Rh9C10L3QuNC1IOKEliAxNDkvMy8yLzIvMjMg0L7R
giAwMi4wMy4yMDE4DDTQl9Cw0LrQu9GO0YfQtdC90LjQtSDihJYgMTQ5LzcvNi00
NDkg0L7RgiAzMC4xMi4yMDIxMD8GBSqFA2RvBDYMNNCf0JDQmtCcIMKr0JrRgNC4
0L/RgtC+0J/RgNC+IEhTTcK7INCy0LXRgNGB0LjQuCAyLjAwDAYFKoUDZHIEAwIB
ADBDBgNVHSAEPDA6MAgGBiqFA2RxATAIBgYqhQNkcQIwCAYGKoUDZHEDMAgGBiqF
A2RxBDAIBgYqhQNkcQUwBgYEVR0gADAOBgNVHQ8BAf8EBAMCAQYwDwYDVR0TAQH/
BAUwAwEB/zAdBgNVHQ4EFgQUyRNYsUynYjp+0j88pucUfJ1wo4YwCgYIKoUDBwEB
AwIDQQCCSXhICg3SZ/TTCtRJpBFuXGSy3PeZTEeOwdOHIv0tWiN2q0mPRzB/o6r9
MXjGqdzfYGtCrq1l5FsXZOI5c/2S
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIIFFTCCBMKgAwIBAgIRAOrKMs71+XnWjTxOTyzGh6QwCgYIKoUDBwEBAwIwggEk
MR4wHAYJKoZIhvcNAQkBFg9kaXRAbWluc3Z5YXoucnUxCzAJBgNVBAYTAlJVMRgw
FgYDVQQIDA83NyDQnNC+0YHQutCy0LAxGTAXBgNVBAcMENCzLiDQnNC+0YHQutCy
0LAxLjAsBgNVBAkMJdGD0LvQuNGG0LAg0KLQstC10YDRgdC60LDRjywg0LTQvtC8
IDcxLDAqBgNVBAoMI9Cc0LjQvdC60L7QvNGB0LLRj9C30Ywg0KDQvtGB0YHQuNC4
MRgwFgYFKoUDZAESDTEwNDc3MDIwMjY3MDExGjAYBggqhQMDgQMBARIMMDA3NzEw
NDc0Mzc1MSwwKgYDVQQDDCPQnNC40L3QutC+0LzRgdCy0Y/Qt9GMINCg0L7RgdGB
0LjQuDAeFw0yMTA3MDIxMjQxNDdaFw0zOTA3MDIxMjQxNDdaMIIBJDEeMBwGCSqG
SIb3DQEJARYPZGl0QG1pbnN2eWF6LnJ1MQswCQYDVQQGEwJSVTEYMBYGA1UECAwP
Nzcg0JzQvtGB0LrQstCwMRkwFwYDVQQHDBDQsy4g0JzQvtGB0LrQstCwMS4wLAYD
VQQJDCXRg9C70LjRhtCwINCi0LLQtdGA0YHQutCw0Y8sINC00L7QvCA3MSwwKgYD
VQQKDCPQnNC40L3QutC+0LzRgdCy0Y/Qt9GMINCg0L7RgdGB0LjQuDEYMBYGBSqF
A2QBEg0xMDQ3NzAyMDI2NzAxMRowGAYIKoUDA4EDAQESDDAwNzcxMDQ3NDM3NTEs
MCoGA1UEAwwj0JzQuNC90LrQvtC80YHQstGP0LfRjCDQoNC+0YHRgdC40LgwZjAf
BggqhQMHAQEBATATBgcqhQMCAiMBBggqhQMHAQECAgNDAARAznwizLAQgEZQecg2
Re7+vfP0h4I5BVHNCzC3UQsCNZdAxM2nn/pKJGaDI8hv4DM3FF3reU2jWDO7Sp66
fTrSZqOCAcIwggG+MIH1BgUqhQNkcASB6zCB6Aw00J/QkNCa0JwgwqvQmtGA0LjQ
v9GC0L7Qn9GA0L4gSFNNwrsg0LLQtdGA0YHQuNC4IDIuMAxD0J/QkNCaIMKr0JPQ
vtC70L7QstC90L7QuSDRg9C00L7RgdGC0L7QstC10YDRj9GO0YnQuNC5INGG0LXQ
vdGC0YDCuww10JfQsNC60LvRjtGH0LXQvdC40LUg4oSWIDE0OS8zLzIvMi8yMyDQ
vtGCIDAyLjAzLjIwMTgMNNCX0LDQutC70Y7Rh9C10L3QuNC1IOKEliAxNDkvNy82
LzEwNSDQvtGCIDI3LjA2LjIwMTgwPwYFKoUDZG8ENgw00J/QkNCa0JwgwqvQmtGA
0LjQv9GC0L7Qn9GA0L4gSFNNwrsg0LLQtdGA0YHQuNC4IDIuMDBDBgNVHSAEPDA6
MAgGBiqFA2RxATAIBgYqhQNkcQIwCAYGKoUDZHEDMAgGBiqFA2RxBDAIBgYqhQNk
cQUwBgYEVR0gADAOBgNVHQ8BAf8EBAMCAQYwDwYDVR0TAQH/BAUwAwEB/zAdBgNV
HQ4EFgQUGXePu4Lmb8hfk/AVHZMiodatDCYwCgYIKoUDBwEBAwIDQQCFR821LAew
8tY3/fYGGXQNN5g6E74pjdurOqOoZbS+0WjiGsLPVzcrAts2l7BRGowjZ3MVAmHJ
L/NjPYAClFu0
-----END CERTIFICATE-----
`

type AttrType int

const (
	GivenName AttrType = AttrType(NID_givenName)
	Surname   AttrType = AttrType(NID_surname)
	SNILS     AttrType = AttrType(NID_SNILS)
	INN       AttrType = AttrType(NID_INN)
	CN        AttrType = AttrType(NID_CN)
)

type CAType int

const (
	QualifiedCA CAType = iota
	NotQualifiedCA
)

type Pkcs7 struct {
	data []byte
}

func FromBytes(bytes []byte) *Pkcs7 {
	return &Pkcs7{
		data: bytes,
	}
}

func (p *Pkcs7) Verify(ca CAType, data []byte) bool {
	if ca == QualifiedCA {
		return VerifyPkcs7([]byte(qualified_ca), p.data, data)
	} else {
		return VerifyPkcs7([]byte(not_qualified_ca), p.data, data)
	}
}

func (p *Pkcs7) GetAttr(attrType AttrType) string {
	return GetPkcs7Attr(p.data, int(attrType))
}