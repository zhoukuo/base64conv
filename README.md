# base64conv —— base64文件编解码工具
[![Maintainability](https://api.codeclimate.com/v1/badges/62fc1698aec14ce44d9c/maintainability)](https://codeclimate.com/github/zhoukuo/base64conv/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/62fc1698aec14ce44d9c/test_coverage)](https://codeclimate.com/github/zhoukuo/base64conv/test_coverage)
## Base64简介

Base64是网络上最常见的用于传输8Bit字节码的编码方式之一，Base64就是一种基于64个可打印字符来表示二进制数据的方法。可查看RFC2045～RFC2049，上面有MIME的详细规范。

Base64编码是从二进制到字符的过程，可用于在HTTP环境下传递较长的标识信息。采用Base64编码具有不可读性，需要解码后才能阅读。

Base64由于以上优点被广泛应用于计算机的各个领域，然而由于输出内容中包括两个以上“符号类”字符（+, /, =)，不同的应用场景又分别研制了Base64的各种“变种”。为统一和规范化Base64的输出，Base62x被视为无符号化的改进版本。


## base64conv工具说明

base64conv 提供base64文件（图片、PDF）的编解码功能。

具体说，当传入base64编码的文件时，会对base64文件解码，然后写入到output指定的文件中；
相反，如果传入为非base64编码的文件，程序会对文件进行base64编码，并写入到output指定的文件中。

在命令行下执行: base64conv.exe -input <文件名> -output <文件名>


周阔
2018-11-28
