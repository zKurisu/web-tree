> Add

```sh
web-tree add --tree=baba/lala --node --link=https://orkarin.com,https://kurisu.com --alias=haha,lala --desc=desc1,desc2 --label=label1,label2
web-tree add --tree=hello --node --link=https://test1.com,https://test2.com --alias=haha,lala --desc=desc1,desc2 --label=label1,label2
```

> Del

```sh
web-tree del --tree=baba/lala --node --link=https://orkarin.com
```

> Edit

```sh
web-tree edit --tree=hello --node --link=https://test1.com --nalias=haha,lala --ndesc=descla,descha --nlabel=labella,labelha
```

> Move

```sh
web-tree move --tree=baba hello
```

```sh
web-tree move --tree=hello --node --link=https://test1.com balabala
```

