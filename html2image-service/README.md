# 基于 WkhtmltoX 的html 转 image

服务器需要安装 WkhtmltoX 才能够使用这个

linux 下载 地址  wget https://github.com/wkhtmltopdf/wkhtmltopdf/releases/download/0.12.4/wkhtmltox-0.12.4_linux-generic-amd64.tar.xz 

unxz wkhtmltoimage-0.11.0_rc1-static-amd64.tar.bz2

tar -xvf wkhtmltox-0.12.4_linux-generic-amd64.tar.xz

mv wkhtmltox/bin/* /usr/local/bin/

rm -rf wkhtmltox

rm -f wkhtmltox-0.12.4_linux-generic-amd64.tar.xz

安装依赖 Cent OS 7

yum install -y libpng	
	
yum install -y libjpeg
	
yum install -y openssl
	
yum install -y icu
	
yum install -y libX11
	
yum install -y libXext
	
yum install -y libXrender
	
yum install -y xorg-x11-fonts-Type1
	
yum install -y xorg-x11-fonts-75dpi】

参考 https://gist.github.com/paulsturgess/cfe1a59c7c03f1504c879d45787699f5


Mac 下可以使用 HomeBrew 安装

