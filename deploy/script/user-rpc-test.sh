repo_addr='crpi-j6xgprtkk3jhxlja.us-west-1.personal.cr.aliyuncs.com/chat-qy/user-rpc-dev'
tag='latest'

container_name="chat-user-rpc-test"
docker stop ${container_name}
docker rm ${container_name}
docker rmi ${repo_addr}:${tag}
docker pull ${repo_addr}:${tag}
docker run -p 10000:10000 --name=${container_name} -d ${repo_addr}:${tag} # 为什么不用指定网络也可以与compose中的中间件成功？？