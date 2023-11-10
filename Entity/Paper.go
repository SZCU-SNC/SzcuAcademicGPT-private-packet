package Entity

type Paper struct {
	Title    string
	Headings []HeadingContent
	Authors  []string
	Keywords []string
}

type HeadingContent struct {
	Title       string
	Content     string
	Subheadings []HeadingContent
}

//	e.g.
// paper := Paper{
//     Title: "论文标题",
//     Headings: []HeadingContent{
//         {
//             Title: "一级标题",
//             Content: "一级标题的正文...",
//             Subheadings: []HeadingContent{
//                 {
//                     Title: "二级标题1",
//                     Content: "二级标题1的正文...",
//                     Subheadings: nil, // 可以继续添加更深层级的标题和正文
//                 },
//                 {
//                     Title: "二级标题2",
//                     Content: "二级标题2的正文...",
//                     Subheadings: nil,
//                 },
//             },
//         },
//     },
//     Authors: []string{"作者1", "作者2"},
//     Keywords: []string{"关键词1", "关键词2"},
// }
