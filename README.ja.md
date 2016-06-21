unmake - Documentation builder for Makefile
===========================================

unmakeはMakefile用のドキュメンテーションツールです。
Makefileのセクションにコメントを埋め込むことでusageとして利用することができます。

Install
-------

Linux or OS X

```
$ curl -L https://github.com/TakesxiSximada/unmake/blob/develop/bootstrap.sh -o bootstrap.sh
$ sh bootstrap.sh
```

How to use It
-------------

Insert next line in your root Makefile top.

```
UNMAKE_CURRENT_MAKEFILE := Makefile
include unmake/Makefile
```

`UNMAKE_CURRENT_MAKEFILE` は最上位のMakefileへのPATHです。
通常は `Makefile` としておけば問題ないでしょう。
`include unmake/Makefile` はunmake/Makefileをincludeしてunmakeのコマンドを取り込んでいます。

Syntax
------

```
UNMAKE_CURRENT_MAKEFILE := Makefile
include unmake/Makefile

.PHONY: example
example:
    @## target=ARGET
    @# exampleのcommandです。
    @# echo "OK" のみ実行します。

    echo "OK"
```

`.PHONY` 指定は必ずセクションの直前に配置する必要があります。
セクション内の `@#` で続くコメントがusageになります。
一行目の `@##` は受け付ける引数などを記述します。

上記のMakefileは以下のような表示になります。

```
$ make help
* example target=ARGET

      exampleのcommandです。
      echo "OK" のみ実行します。


* help

      Display usage.


* unmake-sync

      Sync unmake


* unmake-save

      Save unmake


* unmake-uninstall

      Remove unmake
```

以下のコマンドはunmakeが追加しています。

* help
    usageを表示します。
* unmake-sync
    unmakeの最新のインストールします。
* unmake-save
    unmakeをgitでcommitします。
* unmake-uninstall
    unmakeを削除します。
