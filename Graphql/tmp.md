~~~
xorm reverse mysql "root:123456@(82.156.15.16:3306)/test" $GOPATH/pkg/mod/github.com/go-xorm/cmd/xorm@v0.0.0-20190426080617-f87981e709a1/templates/goxorm  /Users/guomiao/Desktop/go/API/Graphql/app/dao
~~~



# 定义数据类型Article
type Article{
id:ID!
text:String
isPublished:Boolean
author:User!
comments:[Comment!]
}

type User{
id:ID
name:String
}

type Comment{
id:ID!
desc:String
author:User!
}
input NewArticle {
text: String!
authorId: ID!
}
type Query{
articles:[Article!]!
article(id:Int):Article!
}
type Mutation{
createArticle(input:NewArticle!):Article!
updateArticle(id:Int):Article!
deleteArticle(id:Int):Article!
}
