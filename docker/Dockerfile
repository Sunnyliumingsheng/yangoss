# 使用官方 nginx 镜像作为基础镜像
FROM nginx:latest

# 创建目录 /img，挂载图片
RUN mkdir -p /img

# 将本地的 nginx.conf 配置文件复制到镜像中
COPY nginx.conf /etc/nginx/nginx.conf

# 将本地的 default.conf 配置文件复制到镜像中
COPY default.conf /etc/nginx/conf.d/default.conf

# 将本地的图片目录挂载到 /img 目录
VOLUME ["/img"]

# 暴露 80 端口
EXPOSE 80

# 启动 nginx
CMD ["nginx", "-g", "daemon off;"]
