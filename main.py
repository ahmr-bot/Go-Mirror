# main.py
# By GitHub @kingc2022

import json
import os

import requests

cores = ["Akarin","Bedrock_Server","BungeeCord","CraftBukkit","HexaCord","Hose","MITE","Minecraft_Server","Mohist","NukkitX","PaperSpigot","PocketMine-MP","Spigot","SpongeForge","SpongeVanilla","Velocity","WaterFall"]


def download(core, name, url):
    url = url.replace("servers", "server")
    url = url.replace("cdn", "prdx")
    print(f"正在下载 {url} : {name}")
    response = requests.get(url)
    with open(f"root/{core}/{name}", 'wb') as f:
        f.write(response.content)
    print(f"已下载 root/{core}/{name}")


def check():
    global cores
    if not os.path.exists("root"):
        os.mkdir("root")
    os.chdir("root")
    for i in range(0, len(cores)):
        if not os.path.exists(f"{cores[i]}"):
            print(f"检测到 root/{cores[i]} 文件夹缺失，自动创建中...")
            os.mkdir(cores[i])
    os.chdir("..")


def main():
    global cores
    for i in range(0, len(cores)):
        response = requests.get(f"https://mirror.zerodream.net/?action=getlist&version={cores[i]}")
        response.encoding = response.apparent_encoding
        response_json = json.loads(response.text)
        for j in range(0, len(response_json)):
            if os.path.exists(f"root/{cores[i]}/{response_json[j]['name']}"):
                print(f"检测到 root/{cores[i]}/{response_json[j]['name']} 已存在，自动跳过")
                continue
            else:
                download(cores[i], response_json[j]["name"], response_json[j]["file"])


if __name__ == '__main__':
    check()
    main()
