# GitHub Pages + Hugo 建博客站


借助于 Github Pages 提供的静态网站托管服务，并采用了 Hugo 这一开源项目，加快了建站流程，而且有多种开源网站主题可供选择.

<!--more-->

## 前言

在博客网站搭建完成之后，有充分的理由相信，自己在未来很长一段时间内将不会再次重复建站。

**常言道天有不测风云，为了防止各种意外情况发生，导致本博客网站无法正常使用，同时防止自己忘记搭建流程，故记录于此。**

## 效果

![效果图](/img/Kyten-blog-outline.png)

## 相关知识简介

### Github Pages

GitHub Pages 是一个免费的静态网站托管服务，它允许用户通过 GitHub 存储库来托管和发布网页，可以使用它来展示项目文档、博客或个人简历。

![github pages](/img/github-pages-intro.png)

现阶段，Github Pages 支持公共存储库的免费的托管；对于私有仓库，需要进行缴费。

### Hugo

官方号称，[Hugo](https://gohugo.io/) 是世界上最快的网站建设框架(The world’s fastest framework for building websites)。

![Hugo](/img/hugo-intro.png)

## Steps

### Github 仓库创建

需要创建两个仓库，一个用于网站源码管理(`sA`)，一个用于网站部署(`sB`):

- `sA` 可以是 `public`，也可以是 `private`；
- `sB` 仓库的名称必须是 `username.github.io`（`username` 是 Github `Accout` 中`username`，不是 `profile` 中的 `Name`），同时还需要添加 `README.md`；

### 使用 Hugo 创建网站

首先，使用 Git 将 `sA` 拉取下来:

```bash
~/ $ git clone https://github.com/lutianen/kyten-blog.git
```

然后，进入本地的 `sA` 目录（即，`kyten-blog`）下，使用 hugo 建站：

```bash
# Linux: Install
~/kyten-blog $ sudo pacman -S hugo
~/kyten-blog $ hugo version 

# 建站，然后将生成的内容复制到 `sA` 仓库中
~/kyten-blog $ hugo new kyten-blog
~/kyten-blog $ mv kyten-blog/ .
~/kyten-blog $ rm kyten-blog -rf
```

### Hugo 设置网站主题

可以从 [Hugo Themes](https://themes.gohugo.io/) 挑选合适的主题进行应用：

```bash
~/kyten-blog $ cd themes
~/kyten-blog/themes $ git clone https://github.com/kakawait/hugo-tranquilpeak-theme.git tranquilpeak
```

安装 Hugo 主题后，根据个人情况修改相应的配置文件即可；

### 文章管理

#### 启动 Hugo server

启动本地 server：

```bash
~/kyten-blog $ hugo server -D
```

浏览器打开 [http://localhost:1313/](http://localhost:1313/) 进行预览；

#### 新建文章

```bash
~/kyten-blog $ hugo new content `post/Golang/Go.md` # `post/Golang/Go.md` 表明 markdown 的路径
```

#### 部署文章

##### 构建 Hugo 网站相关静态文件

Hugo 将构建完成的静态内容保存到 `sA` 仓库中的 `public` 文件夹中；

```bash
~/kyten-blog $ hugo
```

##### 部署

进入 `public` 目录，利用 Git 进行管理该文件夹，并推送到远程 `sB` 仓库中：

```bash
~/kyten-blog/public $ git init
~/kyten-blog/public $ git commit -m "first commit"
~/kyten-blog/public $ git branch -M master
~/kyten-blog/public $ git remote add origin https://github.com/lutianen/test.git
~/kyten-blog/public $ git push -u origin master
```

自动化部署：`deploy.sh`

```bash
#!/bin/bash 

echo -e "\033[0;32mDeploying updates to GitHub...\033[0m" 

# Build the project. 
hugo # if using a theme, replace with hugo -t 

# Go To Public folder 
cd public 
# Add changes to git. 
git add . 

# Commit changes. 
msg="rebuilding site `date` " 

echo -e "\033[0;32m$msg\033[0m"

if [ $# -eq 1 ] 
    then msg="$1" 
fi 

git commit -m "$msg" 
# Push source and build repos. 
git push origin master 

# Come Back up to the Project Root 
cd ..


```

#### 删除文章

进入 `kyten-blog/post/` 目录中，删除，目标文件夹（包含相关文章资源）即可；

NOTE：`kyten-blog/public` 中相关文件可以删除，也可以不删除，推荐删除；

### Google Analytics（分析）

首先，在 [GA](https://analytics.google.com/) 网站中注册、设置完成相应选项，即可获取 ID：`G-XXXXXXXXXX`；
然后在 `layout/_default/baseof.html` 文件中添加以下代码即可：

```HTML
<!-- Google tag (gtag.js) -->
<script async src="https://www.googletagmanager.com/gtag/js?id=G-XXXXXXXXXX"></script>
<script>
window.dataLayer = window.dataLayer || [];
function gtag(){dataLayer.push(arguments);}
gtag('js', new Date());

gtag('config', 'G-XXXXXXXXXX');
</script>
```

该段代码获取方法如下：[Google Analytics（分析）](https://www.google.com/analytics/web) -> **管理** -> "媒体资源设置"列中的**数据流** -> 网站 -> 对应的数据流 -> "Google 代码" 下的**查看代码说明** -> "安装说明" 选择**手动添加**

> **建议添加完成后，在GA分析中进行测试，确保生效**

### 评论系统 - Gitalk

1. 申请注册[新应用](https://github.com/settings/applications/new)，并填写以下相关内容：

    - `Application name`: 随意
    - `Homepage URL`: 包含 `http/https` 前缀，例如`https://lutianen.github.io`
    - `Authorization callback URL`: 和上面 `Homepage URL` 保持一致就行

    ![github application](/img/github-application.png)

2. 注册完成后，手动生成 `Client secrets`(*只会出现一次*)，同时获得 `Client ID`.

3. 最后，在主题设置中填写相应信息即可，例如 `LoveIt` 中的 `config.toml`:

    ```TOML
    [params.page.comment]
        enable = true

        [params.page.comment.gitalk]
        enable = true
        owner = "lutianen"
        repo = "gitalk"
        clientId = "xxxxxxxxxxxxxxxxxxxx"
        clientSecret = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
    ```

## Problem And Solution

### 添加图片不显示

Hugo 的配置文件和文章中的引用图片都是以 static 作为根目录，因此图片无法显示的解决方案如下：

1. 将图片放入 `static/img` 目录下
2. 在文章中的图片引用方式为：`/img/xxx.png`
3. 无法采用 Typora 等软件进行预览，需要在网页中进行预览: [http://localhost:1313/](http://localhost:1313/)

## References

- [Abot Github Pages](https://docs.github.com/en/pages/getting-started-with-github-pages/about-github-pages)
- [Hugo](https://gohugo.io/)
- [参考文章](https://zz2summer.github.io/github-pages-hugo-%E6%90%AD%E5%BB%BA%E4%B8%AA%E4%BA%BA%E5%8D%9A%E5%AE%A2)
- [Gitalk 评论系统安装](https://www.gagahappy.com/gitalk-install/)

