/*
 * @Author MierX
 * @Title 10正则表达式3
 * @Date 2023-03-27 周一
 * @Time 13:34:16
 */

package main

import (
	"fmt"
	"regexp"
)

func main() {
	//反引号 原生字符串
	buf := `
<!DOCTYPE html>
<html>
 <head> 
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" /> 
<meta name="description" content="Golang语言社区 维护的Go语言标准库文档，Go语言学习者的首选中文文档，Golang语言社区网址：www.Golang.Ltd" />
<meta name="keywords" content="Go，Go语言，Golang语言，Golang语言社区，www.Golang.Ltd，彬哥和Go语言，Go语言标准库中文文档" />
  <title>Go库文档</title> 
  <link type="text/css" rel="stylesheet" href="./index_files/style.css" /> 
  <link type="text/css" rel="stylesheet" href="./doc/Goguifan/style.css">
  <link rel="stylesheet" href="./doc/Goguifan/jquery.treeview.css">
  
  <style type="text/css"></style> 
  <style>@-moz-keyframes nodeInserted{from{opacity:0.99;}to{opacity:1;}}@-webkit-keyframes nodeInserted{from{opacity:0.99;}to{opacity:1;}}@-o-keyframes nodeInserted{from{opacity:0.99;}to{opacity:1;}}@keyframes nodeInserted{from{opacity:0.99;}to{opacity:1;}}embed,object{animation-duration:.001s;-ms-animation-duration:.001s;-moz-animation-duration:.001s;-webkit-animation-duration:.001s;-o-animation-duration:.001s;animation-name:nodeInserted;-ms-animation-name:nodeInserted;-moz-animation-name:nodeInserted;-webkit-animation-name:nodeInserted;-o-animation-name:nodeInserted;}</style>
  <style type="text/css" id="8726384429"></style>
 </head> 
 
 <body>
 <div id="box"></div>
 <script type="text/javascript">
     function setStorage(){
         if(localStorage.counter){
             localStorage.counter=Number(localStorage.counter)+1;
         }else{
             localStorage.counter=1;
         }
         return localStorage.counter;
     }
     var counter=setStorage();
     var oBox=document.getElementById('box');
	var xhr=new XMLHttpRequest();
	//配置Ajax对象
	xhr.open('get','http://47.107.125.75:8867/num?');
	//发送请求
	xhr.send();
	//获取服务器端响应的数据
	xhr.onload=function(){
	//console.log(xhr.responseText.length)
	if (xhr.responseText.length == 19 ) {
	   oBox.innerHTML="【Golang语言情怀】网站被访问的次数为<b>"+counter+"</b>次";
	}else{
	   oBox.innerHTML="【Golang语言情怀】网站被访问的次数为<b>"+xhr.responseText+"</b>次";
	}
	}
 </script>
 </body>
 
 <body > 
  <div id="lowframe" style="position: fixed; bottom: 0; left: 0; height: 0; width: 100%; border-top: thin solid grey; background-color: white; overflow: auto;">
    ... 
  </div>
  <!-- #lowframe --> 
  <div id="page" class="wide" tabindex="-1" style="outline: 0px;"> 
   <div class="container"> 
 
   <script async="" src="https://www.google-analytics.com/analytics.js"></script><script type="text/javascript">window.initFuncs = [];</script>
   </head>
   <body>
   
   <div id="lowframe" style="position: fixed; bottom: 0; left: 0; height: 0; width: 100%; border-top: thin solid grey; background-color: white; overflow: auto;">
   ...
   </div><!-- #lowframe -->
   
   <div id="topbar" class="wide"><div class="container">
   <div id="menu" style="min-width: 650px;">
   <a href="./doc/Goneicun.htm">Go内存模型</a>
   <a href="./doc/Goguifan.htm">编程规范</a>
   <a href="./doc/Gomingling.htm">Go命名</a>
   <a href="./doc/RHGo.htm">如何使用Go编程</a>
   <a href="./doc/GoSX.htm">Go实效编程</a>
   <!--
     <a id="playgroundButton" href="http://play.golang.org/" title="显示 Go 操场" style="display: inline;">运行</a>
   -->
   </div>
   <div id="heading"><a href="#">微信公众账号：<b>Golang语言情怀</b>（原:Golang语言社区）</a></div>
   </form>
   </div></div>
   
   
   <div id="playground" class="play">
   	<div class="input"><textarea class="code">package main
   
   import "fmt"
   
   func main() {
   	fmt.Println("Hello, 世界")
   }</textarea></div>
   	<div class="output"></div>
   	<div class="buttons">
   		<a class="run" title="运行此代码[Shift-Enter]">运行</a>
   		<a class="fmt" title="格式化此代码">格式化</a>
   		<a class="share" title="分享此代码">分享</a>
   	</div>
   </div>
   <div id="page" class="wide" tabindex="-1" style="outline: 0px;">
   <div class="container">
   
    <h1>Golang标准库文档</h1> 
<hr />
<a href="https://golang.google.cn/dl/" target="_blank">Ⅰ.Golang语言标准库下载(官方地址)</a> </br>
<a href="https://godoc.org/" target="_blank">Ⅱ.Golang语言第三方包下载</a>  </br>
<a href="http://liteide.org/cn/" target="_blank">Ⅲ.Golang语言编辑器(☯liteide)下载</a>  </br>
<a href="http://ol.golang.ltd" target="_blank">IV.Golang语言在线编辑器(随时随地玩转Go语言)</a> </br>
<hr />
<a href="https://pan.baidu.com/s/1AndukeGoBHnz7cYIDVUBdA" target="_blank">Ⅰ.服务器系统下载(仅提供:centos;提取码:game)</a> </br>
<a href="https://dev.mysql.com/downloads/mysql/" target="_blank">Ⅱ.数据库下载(仅提供:mysql)</a>  </br>
<a href="https://pan.baidu.com/s/1ZAMAwB_DjZSvjjW4peW67w" target="_blank">Ⅲ.虚拟化工具下载(提取码:game)</a>  </br>
<hr /></br>

<form action="s.php" method="get">
<input type="text" name="s"><input type="submit" value="搜索">
</form>
   
	<h2 id="other">★分布式游戏服务器框架★</h2>
	<h3 id="subrepo">LollipopGo分布式游戏框架</h3>
	<p>LollipopGo框架由Golang语言社区发起人彬哥维护开发，框架v2.8.x商业化版本已经更新，同时v3.0.x框架支持3D游戏服务器引擎，
	包括寻路引起，物理碰撞等
	<a href="/doc/go1compat">物理引擎</a>相关。
	可通过“<a href="/cmd/go/#hdr-Download_and_install_packages_and_dependencies" target="_blank">go mod</a>”安装，代码库的<a href="http://godoc.org/-/subrepo">文档</a>和<a href="https://github.com/golang">源码</a>可通过相应的链接访问</p>
	<ul> 
	 <li><a href="https://github.com/Golangltd/LollipopGo" target="_blank">框架GitHub地址</a> — 期待Star。</li> 
	 <li><a href="https://space.bilibili.com/389368547?from=search&seid=4761602022286063108" target="_blank">框架B站视频</a> — 彬哥讲解。</li> 
	 <li><a href="https://gopher.ke.qq.com" target="_blank">腾讯课堂视频</a> — 彬哥讲解。</li> 
	 <li><a href="https://www.bilibili.com/video/BV1aE411H7W6" target="_blank">斗兽棋第一期实战项目</a> — LollipopGo2.4版本。</li>
	 </ul> 
	 <hr />
	 <h2 id="other">Leaf游戏服务器框架拓展</h2>
	 <h3 id="subrepo">Leaf游戏服务器框架拓展基础库Nest</h3>
	 <p>Nest框架是彬哥之前项目中的基础库，对Leaf框架进行了拓展，消息Id自定义等。获取最新信息请关注公众账号：Golang语言游戏服务器！
	 <ul> 
	  <li><a href="https://github.com/Golangltd/nest" target="_blank">Nest框架GitHub地址</a> — 期待Star。</li> 
	  </ul> 
	 
<hr />
    <div id="nav"></div> 
    <h2 id="pkg-subdirectories">子目录</h2> 
    <div id="manual-nav"> 
     <dl> 
      <dt>
       <a href="#stdlib">标准库</a>
      </dt> 
      <dt>
       <a href="#other">其它包</a>
      </dt> 
      <dd>
       <a href="#subrepo">子代码库</a>
      </dd> 
      <dd>
       <a href="#community">社区</a>
      </dd> 
     </dl> 
    </div>
    <h2 id="stdlib">标准库</h2> 
    <img class="gopher" src="./index_files/pkg.png" /> 
    <table class="dir"> 
     <tbody>
      <tr> 
       <th>名称</th> 
       <th>&nbsp;&nbsp;&nbsp;&nbsp;</th> 
       <th style="text-align: left; width: auto">摘要</th> 
      </tr> 
      <tr> 
       <td class="name"><a id="archive" href="#archive">archive</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto"></td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/archive_tar.htm">tar</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">tar包实现了tar格式压缩文件的存取.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/archive_zip.htm">zip</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">zip包提供了zip档案文件的读写服务.</td> 
      </tr> 
      <tr> 
       <td class="name"><a href="./pkg/bufio.htm">bufio</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">bufio 包实现了带缓存的I/O操作.</td> 
      </tr> 
      <tr> 
       <td class="name"><a href="./pkg/builtin.htm">builtin</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">builtin 包为Go的预声明标识符提供了文档.</td> 
      </tr> 
      <tr> 
       <td class="name"><a href="./pkg/bytes.htm">bytes</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">bytes包实现了操作[]byte的常用函数.</td> 
      </tr> 
      <tr> 
       <td class="name"><a id="compress" href="#compress">compress</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto"></td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/compress_bzip2.htm">bzip2</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">bzip2包实现bzip2的解压缩.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/compress_flate.htm">flate</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">flate包实现了deflate压缩数据格式，参见RFC 1951.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/compress_gzip.htm">gzip</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">gzip包实现了gzip格式压缩文件的读写，参见RFC 1952.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/compress_lzw.htm">lzw</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">lzw包实现了Lempel-Ziv-Welch数据压缩格式，这是一种T. A. Welch在“A Technique for High-Performance Data Compression”一文（Computer, 17(6) (June 1984), pp 8-19）提出的一种压缩格式.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/compress_zlib.htm">zlib</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">zlib包实现了对zlib格式压缩数据的读写，参见RFC 1950.</td> 
      </tr> 
      <tr> 
       <td class="name"><a id="container" href="#container">container</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto"></td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/container_heap.htm">heap</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">heap包提供了对任意类型（实现了heap.Interface接口）的堆操作.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/container_list.htm">list</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">list包实现了双向链表.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/container_ring.htm">ring</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">ring实现了环形链表的操作.</td> 
      </tr>
      <tr> 
       <td class="name"><a id="context" href="./pkg/context.htm">context</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">Package context defines the Context type, which carries deadlines, cancelation signals, and other request-scoped values across API boundaries and between processes.</td> 
      </tr>
      <tr> 
       <td class="name"><a id="crypto" href="./pkg/crypto.htm">crypto</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">crypto包搜集了常用的密码（算法）常量.</td> 
      </tr>
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/crypto_aes.htm">aes</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">aes包实现了AES加密算法，参见U.S. Federal Information Processing Standards Publication 197.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/crypto_cipher.htm">cipher</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">cipher包实现了多个标准的用于包装底层块加密算法的加密算法实现.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/crypto_des.htm">des</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">des包实现了DES标准和TDEA算法，参见U.S. Federal Information Processing Standards Publication 46-3.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/crypto_dsa.htm">dsa</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">dsa包实现FIPS 186-3定义的数字签名算法（Digital Signature Algorithm），即DSA算法.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/crypto_ecdsa.htm">ecdsa</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">ecdsa包实现了椭圆曲线数字签名算法，参见FIPS 186-3.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/crypto_elliptic.htm">elliptic</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">elliptic包实现了几条覆盖素数有限域的标准椭圆曲线.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/crypto_hmac.htm">hmac</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">hmac包实现了U.S. Federal Information Processing Standards Publication 198规定的HMAC（加密哈希信息认证码）.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/crypto_md5.htm">md5</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">md5包实现了MD5哈希算法，参见RFC 1321.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/crypto_rand.htm">rand</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">rand包实现了用于加解密的更安全的随机数生成器.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/crypto_rc4.htm">rc4</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">rc4包实现了RC4加密算法，参见Bruce Schneier's Applied Cryptography.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/crypto_rsa.htm">rsa</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">rsa包实现了PKCS#1规定的RSA加密算法.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/crypto_sha1.htm">sha1</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">sha1包实现了SHA1哈希算法，参见RFC 3174.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/crypto_sha256.htm">sha256</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">sha256包实现了SHA224和SHA256哈希算法，参见FIPS 180-4.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/crypto_sha512.htm">sha512</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">sha512包实现了SHA384和SHA512哈希算法，参见FIPS 180-2.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/crypto_subtle.htm">subtle</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">Package subtle implements functions that are often useful in cryptographic code but require careful thought to use correctly.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/crypto_tls.htm">tls</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">tls包实现了TLS 1.2，细节参见RFC 5246.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/crypto_x509.htm">x509</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">x509包解析X.509编码的证书和密钥.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/crypto_x509_pkix.htm">pkix</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">pkix包提供了共享的、低层次的结构体，用于ASN.1解析和X.509证书、CRL、OCSP的序列化.</td> 
      </tr> 
      <tr> 
       <td class="name"><a href="#database" id="database">database</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto"></td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/database_sql.htm">sql</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">sql 包提供了通用的SQL（或类SQL）数据库接口.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/database_sql_driver.htm">driver</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">driver包定义了应被数据库驱动实现的接口，这些接口会被sql包使用.</td> 
      </tr> 
      <tr> 
       <td class="name"><a id="debug" href="#debug">debug</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto"></td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/debug_dwarf.htm">dwarf</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">Package dwarf provides access to DWARF debugging information loaded from executable files, as defined in the DWARF 2.0 Standard at http://dwarfstd.org/doc/dwarf-2.0.0.pdf</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/debug_elf.htm">elf</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">Package elf implements access to ELF object files.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/debug_gosym.htm">gosym</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">Package gosym implements access to the Go symbol and line number tables embedded in Go binaries generated by the gc compilers.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/debug_macho.htm">macho</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">Package macho implements access to Mach-O object files.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/debug_pe.htm">pe</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">Package pe implements access to PE (Microsoft Windows Portable Executable) files.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/debug_plan9obj.htm">plan9obj</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">Package plan9obj implements access to Plan 9 a.out object files.</td> 
      </tr> 
      <tr> 
       <td class="name"><a id="encoding" href="./pkg/encoding.htm">encoding</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">encoding包定义了供其它包使用的可以将数据在字节水平和文本表示之间转换的接口.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/encoding_ascii85.htm">ascii85</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">ascii85 包是对 ascii85 的数据编码的实现.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/encoding_asn1.htm">asn1</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">asn1包实现了DER编码的ASN.1数据结构的解析，参见ITU-T Rec X.690.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/encoding_base32.htm">base32</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">base32包实现了RFC 4648规定的base32编码.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/encoding_base64.htm">base64</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">base64实现了RFC 4648规定的base64编码.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/encoding_binary.htm">binary</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">binary包实现了简单的数字与字节序列的转换以及变长值的编解码.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/encoding_csv.htm">csv</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">csv读写逗号分隔值（csv）的文件.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/encoding_gob.htm">gob</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">gob包管理gob流——在编码器（发送器）和解码器（接受器）之间交换的binary值.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/encoding_hex.htm">hex</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">hex包实现了16进制字符表示的编解码.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/encoding_json.htm">json</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">json包实现了json对象的编解码，参见RFC 4627.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/encoding_pem.htm">pem</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">pem包实现了PEM数据编码（源自保密增强邮件协议）.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/encoding_xml.htm">xml</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">Package xml implements a simple XML 1.0 parser that understands XML name spaces.</td> 
      </tr> 
      <tr> 
       <td class="name"><a id="errors" href="./pkg/errors.htm">errors</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">error 包实现了用于错误处理的函数.</td> 
      </tr> 
      <tr> 
       <td class="name"><a href="./pkg/expvar.htm">expvar</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">expvar包提供了公共变量的标准接口，如服务的操作计数器.</td> 
      </tr> 
      <tr> 
       <td class="name"><a href="./pkg/flag.htm">flag</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">flag 包实现命令行标签解析.</td> 
      </tr> 
      <tr> 
       <td class="name"><a href="./pkg/fmt.htm">fmt</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">fmt 包实现了格式化I/O函数，类似于C的 printf 和 scanf.</td> 
      </tr> 
      <tr> 
       <td class="name"><a id="go" href="#go">go</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto"></td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/go_ast.htm">ast</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">Package ast declares the types used to represent syntax trees for Go packages.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/go_build.htm">build</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">Package build gathers information about Go packages.</td> 
      </tr>
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/go_constant.htm">constant</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">Package constant implements Values representing untyped Go constants and their corresponding operations.</td>
      </tr>
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/go_doc.htm">doc</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">Package doc extracts source code documentation from a Go AST.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/go_format.htm">format</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">Package format implements standard formatting of Go source.</td> 
      </tr>
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/go_importer.htm">importer</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">Package importer provides access to export data importers.</td>
      </tr>
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/go_parser.htm">parser</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">Package parser implements a parser for Go source files.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/go_printer.htm">printer</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">Package printer implements printing of AST nodes.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/go_scanner.htm">scanner</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">Package scanner implements a scanner for Go source text.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/go_token.htm">token</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">Package token defines constants representing the lexical tokens of the Go programming language and basic operations on tokens (printing, predicates).</td> 
      </tr>
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/go_types.htm">types</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">Package types declares the data types and implements the algorithms for type-checking of Go packages.</td> 
      </tr>
      <tr> 
       <td class="name"><a id="hash" href="#hash">hash</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">hash包提供hash函数的接口.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/hash_adler32.htm">adler32</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">adler32包实现了Adler-32校验和算法，参见RFC 1950.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/hash_crc32.htm">crc32</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">crc32包实现了32位循环冗余校验（CRC-32）的校验和算法.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/hash_crc64.htm">crc64</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">Package crc64 implements the 64-bit cyclic redundancy check, or CRC-64, checksum.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/hash_fnv.htm">fnv</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">fnv包实现了FNV-1和FNV-1a（非加密hash函数）.</td> 
      </tr> 
      <tr> 
       <td class="name"><a href="./pkg/html.htm">html</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">html包提供了用于转义和解转义HTML文本的函数.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/html_template.htm">template</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">template包（html/template）实现了数据驱动的模板，用于生成可对抗代码注入的安全HTML输出.</td> 
      </tr> 
      <tr> 
       <td class="name"><a href="./pkg/image.htm">image</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">image实现了基本的2D图片库.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/image_color.htm">color</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">color 包实现了基本的颜色库。</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/image_color_palette.htm">palette</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">palette包提供了标准的调色板.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/image_draw.htm">draw</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">draw 包提供组装图片的方法.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/image_gif.htm">gif</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">gif 包实现了GIF图片的解码.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/image_jpeg.htm">jpeg</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">jpeg包实现了jpeg格式图像的编解码.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/image_png.htm">png</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">png 包实现了PNG图像的编码和解码.</td> 
      </tr> 
      <tr> 
       <td class="name"><a id="index" href="#index">index</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto"></td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/index_suffixarray.htm">suffixarray</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">suffixarrayb包通过使用内存中的后缀树实现了对数级时间消耗的子字符串搜索.</td> 
      </tr> 
      <tr> 
       <td class="name"><a href="./pkg/io.htm">io</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">io 包为I/O原语提供了基础的接口.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/io_ioutil.htm">ioutil</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">ioutil 实现了一些I/O的工具函数。</td> 
      </tr> 
      <tr> 
       <td class="name"><a href="./pkg/log.htm">log</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">log包实现了简单的日志服务.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/log_syslog.htm">syslog</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">syslog包提供一个简单的系统日志服务的接口.</td> 
      </tr> 
      <tr> 
       <td class="name"><a href="./pkg/math.htm">math</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">math 包提供了基本常数和数学函数。</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/math_big.htm">big</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">big 包实现了（大数的）高精度运算.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/math_cmplx.htm">cmplx</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">cmplx 包为复数提供了基本的常量和数学函数.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/math_rand.htm">rand</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">rand 包实现了伪随机数生成器.</td> 
      </tr> 
      <tr> 
       <td class="name"><a href="./pkg/mime.htm">mime</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">mime实现了MIME的部分规定.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/mime_multipart.htm">multipart</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">multipart实现了MIME的multipart解析，参见RFC 2046.</td> 
      </tr>
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/mime_quotedprintable.htm">quotedprintable</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">Package quotedprintable implements quoted-printable encoding as specified by RFC 2045.</td> 
      </tr>
      <tr> 
       <td class="name"><a href="./pkg/net.htm">net</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">net包提供了可移植的网络I/O接口，包括TCP/IP、UDP、域名解析和Unix域socket.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/net_http.htm">http</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">http包提供了HTTP客户端和服务端的实现.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/net_http_cgi.htm">cgi</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">cgi 包实现了RFC3875协议描述的CGI（公共网关接口）.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/net_http_cookiejar.htm">cookiejar</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">cookiejar包实现了保管在内存中的符合RFC 6265标准的http.CookieJar接口.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/net_http_fcgi.htm">fcgi</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">fcgi 包实现了FastCGI协议.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/net_http_httptest.htm">httptest</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">httptest 包提供HTTP测试的单元工具.</td> 
      </tr>
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/net_http_httptrace.htm">httptrace</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">Package httptrace provides mechanisms to trace the events within HTTP client requests.</td> 
      </tr>
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/net_http_httputil.htm">httputil</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">httputil包提供了HTTP公用函数，是对net/http包的更常见函数的补充.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/net_http_pprof.htm">pprof</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">pprof 包通过提供HTTP服务返回runtime的统计数据，这个数据是以pprof可视化工具规定的返回格式返回的.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/net_mail.htm">mail</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">mail 包实现了解析邮件消息的功能.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/net_rpc.htm">rpc</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">rpc 包提供了一个方法来通过网络或者其他的I/O连接进入对象的外部方法.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/net_rpc_jsonrpc.htm">jsonrpc</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">jsonrpc 包使用了rpc的包实现了一个JSON-RPC的客户端解码器和服务端的解码器.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/net_smtp.htm">smtp</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">smtp包实现了简单邮件传输协议（SMTP），参见RFC 5321.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/net_textproto.htm">textproto</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">textproto实现了对基于文本的请求/回复协议的一般性支持，包括HTTP、NNTP和SMTP.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/net_url.htm">url</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">url包解析URL并实现了查询的逸码，参见RFC 3986.</td> 
      </tr> 
      <tr> 
       <td class="name"><a href="./pkg/os.htm">os</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">os包提供了操作系统函数的不依赖平台的接口.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/os_exec.htm">exec</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">exec包执行外部命令.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/os_signal.htm">signal</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">signal包实现了对输入信号的访问.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/os_user.htm">user</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">user包允许通过名称或ID查询用户帐户.</td> 
      </tr> 
      <tr> 
       <td class="name"><a href="./pkg/path.htm">path</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">path实现了对斜杠分隔的路径的实用操作函数.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/path_filepath.htm">filepath</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">filepath包实现了兼容各操作系统的文件路径的实用操作函数.</td> 
      </tr> 
      <tr>
       <td class="name"><a href="./pkg/plugin.htm">plugin</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">Package plugin implements loading and symbol resolution of Go plugins.</td> 
      </tr>
      <tr>
       <td class="name"><a href="./pkg/reflect.htm">reflect</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">reflect包实现了运行时反射，允许程序操作任意类型的对象.</td> 
      </tr>
      <tr> 
       <td class="name"><a href="./pkg/regexp.htm">regexp</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">regexp包实现了正则表达式搜索.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/regexp_syntax.htm">syntax</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">Package syntax parses regular expressions into parse trees and compiles parse trees into programs.</td> 
      </tr> 
      <tr> 
       <td class="name"><a href="./pkg/runtime.htm">runtime</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">TODO(osc): 需更新 runtime 包含与Go的运行时系统进行交互的操作，例如用于控制Go程的函数.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/runtime_cgo.htm">cgo</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">cgo 包含有 cgo 工具生成的代码的运行时支持.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/runtime_debug.htm">debug</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">debug 包含有程序在运行时调试其自身的功能.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/runtime_pprof.htm">pprof</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">pprof 包按照可视化工具 pprof 所要求的格式写出运行时分析数据.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/runtime_race.htm">race</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">race 包实现了数据竞争检测逻辑.</td> 
      </tr>
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/runtime_trace.htm">trace</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">Go execution tracer.</td> 
      </tr>
      <tr> 
       <td class="name"><a href="./pkg/sort.htm">sort</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">sort 包为切片及用户定义的集合的排序操作提供了原语.</td> 
      </tr> 
      <tr> 
       <td class="name"><a href="./pkg/strconv.htm">strconv</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">strconv包实现了基本数据类型和其字符串表示的相互转换.</td> 
      </tr> 
      <tr> 
       <td class="name"><a href="./pkg/strings.htm">strings</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">strings包实现了用于操作字符的简单函数.</td> 
      </tr> 
      <tr> 
       <td class="name"><a href="./pkg/sync.htm">sync</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">sync 包提供了互斥锁这类的基本的同步原语.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/sync_atomic.htm">atomic</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">atomic 包提供了底层的原子性内存原语，这对于同步算法的实现很有用.</td> 
      </tr> 
      <tr> 
       <td class="name"><a href="./pkg/syscall.htm">syscall</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">Package syscall contains an interface to the low-level operating system primitives.</td> 
      </tr> 
      <tr> 
       <td class="name"><a href="./pkg/testing.htm">testing</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">Package testing provides support for automated testing of Go packages.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/testing_iotest.htm">iotest</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">Package iotest implements Readers and Writers useful mainly for testing.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/testing_quick.htm">quick</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">Package quick implements utility functions to help with black box testing.</td> 
      </tr> 
      <tr> 
       <td class="name"><a id="text" href="#text">text</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto"></td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/text_scanner.htm">scanner</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">scanner包提供对utf-8文本的token扫描服务.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/text_tabwriter.htm">tabwriter</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">tabwriter包实现了写入过滤器（tabwriter.Writer），可以将输入的缩进修正为正确的对齐文本.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/text_template.htm">template</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">template包实现了数据驱动的用于生成文本输出的模板.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/text_template_parse.htm">parse</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">Package parse builds parse trees for templates as defined by text/template and html/template.</td> 
      </tr> 
      <tr> 
       <td class="name"><a href="./pkg/time.htm">time</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">time包提供了时间的显示和测量用的函数.</td> 
      </tr> 
      <tr> 
       <td class="name"><a href="./pkg/unicode.htm">unicode</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">unicode 包提供了一些测试Unicode码点属性的数据和函数.</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/unicode_utf16.htm">utf16</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">utf16 包实现了对UTF-16序列的编码和解码。</td> 
      </tr> 
      <tr> 
       <td class="name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="./pkg/unicode_utf8.htm">utf8</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">utf8 包实现了支持UTF-8文本编码的函数和常量.</td> 
      </tr> 
      <tr> 
       <td class="name"><a href="./pkg/unsafe.htm">unsafe</a></td> 
       <td>&nbsp;&nbsp;&nbsp;&nbsp;</td> 
       <td style="width: auto">unsafe 包含有关于Go程序类型安全的所有操作.</td> 
      </tr> 
     </tbody>
    </table> 
    <h2 id="other">其它包</h2> 
    <h3 id="subrepo">子代码库</h3>
    <p>这些包是 Go 项目的一部分，但并未在主源码树中。它们在比 Go
	核心库更加宽松的<a href="/doc/go1compat">兼容性需求</a>下开发。
	可通过“<a href="/cmd/go/#hdr-Download_and_install_packages_and_dependencies">go get</a>”安装它们，子代码库的<a href="http://godoc.org/-/subrepo">文档</a>和<a href="https://github.com/golang">源码</a>可通过相应的链接访问</p>
    <ul> 
     <li><a href="https://github.com/golang/crypto">crypto</a> — 附加的加密包。</li> 
     <li><a href="https://github.com/golang/image">image</a> — 附加的图像包。</li> 
     <li><a href="https://github.com/golang/net">net</a> — 附加的网络包。</li> 
     <li><a href="https://github.com/golang/sys">sys</a> — 系统调用包。</li> 
     <li><a href="https://github.com/golang/text">text</a> — 文本处理包。</li> 
     <li><a href="https://github.com/golang/tools">tools</a> — godoc、vet、cover 及其它工具。</li> 
     <li><a href="https://github.com/golang/exp">exp</a> — 实验性代码（可能不经警告就更改，请小心对待）。</li> 
    </ul> 
    <h3 id="community">社区</h3> 
    <p> 这些服务可帮你寻找社区提供的开源包。 </p> 
    <ul> 
     <li><a href="http://godoc.org/">GoDoc</a> - 包索引与搜索引擎。</li> 
     <li><a href="http://go-search.org/">Go 搜索</a> - 代码搜索引擎。</li> 
     <li><a href="https://github.com/golang/go/wiki/Projects">Go 维基上的项目</a> - Go 项目策划列表</li>
     <p>备案号：- 蜀ICP备19024073号-2</p> 
    </ul> 
   </div>
   <!-- .container  蜀ICP备18026143号--> 
<div>哈哈哈</div>
<div>哈哈哈2</div>
<div>哈
哈哈
3</div>
  </div>
  <!-- #page --> 
 </body>
</html>
`
	reg := regexp.MustCompile(`<div>(?s:(.*?))</div>`) //提取div内的内容
	if reg == nil {
		fmt.Println("MustCompile err")
		return
	}
	//2、根据规则提取关键信息
	rs := reg.FindAllStringSubmatch(buf, -1)
	//fmt.Println("rs = ", rs)

	//过滤div标签
	for _, text := range rs {
		fmt.Println("text[1] = ", text[1])
	}
}
