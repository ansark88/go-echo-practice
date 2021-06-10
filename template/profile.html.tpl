{{define "profile"}}
<!DOCTYPE html>
<html lang="ja">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="Cache-Control" content="no-cache">
    <title>{{.Title}}</title>
    <link rel="stylesheet" href="static/css/style.css">

    <!-- React関係のライブラリの読み込み -->
    <script src="https://unpkg.com/react@16/umd/react.development.js" crossorigin></script>
    <script src="https://unpkg.com/react-dom@16/umd/react-dom.development.js" crossorigin></script>
    <script src="https://unpkg.com/babel-standalone@6/babel.min.js"></script>

    <!-- 自分で書くスクリプト -->
    <script type="text/babel" src="static/js/profile.jsx" defer></script>
  </head>

  <body>
    <div id="app"></div>
  </body>
</html>
{{end}}