## go语言CGI模式MVC框架小实验

> 学习产物，瞎折腾

* 实现了路由、控制器和视图的功能
* 可以使用cookie

### 使用apache

```<DOCUMENT_ROOT>```为文档根目录，如```/var/www/html```

主配置文件
```html
<VirtualHost *:80>

	ServerAdmin webmaster@localhost
	DocumentRoot "<DOCUMENT_ROOT>"

    DirectoryIndex main
    <Directory "<DOCUMENT_ROOT>">
        Options Indexes FollowSymLinks MultiViews
        AllowOverride All
        RewriteEngine On
        RewriteCond %{REQUEST_FILENAME} !-f
        RewriteCond %{REQUEST_FILENAME} !-d
        RewriteRule . main [L]
        Options +Indexes
    </Directory>

	LogLevel alert rewrite:trace6
	LogLevel info warn

	ErrorLog ${APACHE_LOG_DIR}/error.log
	CustomLog ${APACHE_LOG_DIR}/access.log combined

	Include mods-available/cgi.load
	Include mods-available/cgid.load	
	Include conf-available/serve-cgi-bin.conf

</VirtualHost>
```

cgi配置: ```serve-cgi-bin.conf```
```html
<IfModule mod_alias.c>
	<IfModule mod_cgi.c>
		Define ENABLE_USR_LIB_CGI_BIN
	</IfModule>

	<IfModule mod_cgid.c>
		Define ENABLE_USR_LIB_CGI_BIN
	</IfModule>

	<IfDefine ENABLE_USR_LIB_CGI_BIN>
		ScriptAlias / <DOCUMENT_ROOT>
		<Directory "<DOCUMENT_ROOT>">
			AllowOverride None
			Options +ExecCGI -MultiViews +SymLinksIfOwnerMatch
			Require all granted
                        
		</Directory>
	</IfDefine>
</IfModule>

```

### 环境变量

```bash
IOGO_CONFIG_FILE # 配置文件
IOGO_TEST_ENV    # 测试环境
```