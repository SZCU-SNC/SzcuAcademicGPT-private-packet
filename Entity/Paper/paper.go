package Paper

type Paper struct {
	Title      string
	Authors    []string
	Abstract   string
	Keywords   []string
	Sections   []Section
	References []Reference
	Config     PaperConfig
}

// Section 正文
type Section struct {
	Title       string
	TitleStyle  []string
	SpecialTile []SpecialStyle

	Content        string
	ContentStyle   []string
	SpecialContent []SpecialStyle

	Subsections []Section
	Images      []Image
}

// SpecialStyle 在字符中格式特殊的子字符，比如 aaaaaxxxbbbb 其中xxx加粗
type SpecialStyle struct {
	String string
	index  int
	Style  []string
}

type Reference struct {
	Author          string
	Title           string
	PublicationDate string
}

type Image struct {
	ImageName string
	//跟在第几句话后面
	PositionInPaper int
	Image           []byte
}

type PaperConfig struct {
	//常见的学术写作规范包括APA风格、MLA风格、Chicago风格
	Style string
	//是否双语摘要
	BilingualAbstract bool
	//是否图片按标题来，比如第一张图片在第4章中，图片序号是4.1还是1
	ImagesIndexWithTitle bool
}

/*	e.g.
paper := Paper{
		Title:    "基于人工智能对论文润色",
		Authors:  []string{"张三", "李四"},
		Abstract: "本文通过应用人工智能技术，探讨了如何利用自然语言处理和机器学习技术来改善论文的语言表达和结构，以提高论文的质量和可读性。",
		Keywords: []string{"人工智能", "自然语言处理", "机器学习", "论文润色"},
		Sections: []Section{
			{
				Title: "引言",
				Content: "本章介绍了论文润色的研究背景和意义，以及人工智能技术在论文润色中的应用前景。",
				Subsections: []Section{
					{
						Title:   "研究背景",
						Content: "介绍了当前论文写作中存在的问题，以及提出论文润色的需求和意义。",
					},
					{
						Title:   "人工智能在论文润色中的应用前景",
						Content: "探讨了人工智能技术在自然语言处理和机器学习领域的发展，以及其在论文润色中的潜在应用价值。",
					},
				},
			},
			{
				Title: "相关工作",
				Content: "本章介绍了国内外在论文润色方面的研究现状和已有成果。",
			},
			{
				Title: "人工智能技术在论文润色中的应用",
				Content: "本章详细介绍了如何利用自然语言处理和机器学习技术来改善论文的语言表达和结构，提高论文的质量和可读性。",
			},
		},
		References: []Reference{
			{
				Author:          "Wang, J.",
				Title:           "Improving Writing Quality Using AI",
				PublicationDate: "2023",
			},
			{
				Author:          "Li, H.",
				Title:           "Machine Learning for Natural Language Processing",
				PublicationDate: "2022",
			},
		},
		Config: PaperConfig{
			Style:                "APA",
			BilingualAbstract:    false,
			ImagesIndexWithTitle: true,
		},
	}

 对应的json

{
  "Title": "基于人工智能对论文润色",
  "Authors": ["张三", "李四"],
  "Abstract": "本文通过应用人工智能技术，探讨了如何利用自然语言处理和机器学习技术来改善论文的语言表达和结构，以提高论文的质量和可读性。",
  "Keywords": ["人工智能", "自然语言处理", "机器学习", "论文润色"],
  "Sections": [
    {
      "Title": "引言",
      "Content": "本章介绍了论文润色的研究背景和意义，以及人工智能技术在论文润色中的应用前景。",
      "Subsections": [
        {
          "Title": "研究背景",
          "Content": "介绍了当前论文写作中存在的问题，以及提出论文润色的需求和意义。"
        },
        {
          "Title": "人工智能在论文润色中的应用前景",
          "Content": "探讨了人工智能技术在自然语言处理和机器学习领域的发展，以及其在论文润色中的潜在应用价值。"
        }
      ]
    },
    {
      "Title": "相关工作",
      "Content": "本章介绍了国内外在论文润色方面的研究现状和已有成果。"
    },
    {
      "Title": "人工智能技术在论文润色中的应用",
      "Content": "本章详细介绍了如何利用自然语言处理和机器学习技术来改善论文的语言表达和结构，提高论文的质量和可读性。"
    }
  ],
  "References": [
    {
      "Author": "Wang, J.",
      "Title": "Improving Writing Quality Using AI",
      "PublicationDate": "2023"
    },
    {
      "Author": "Li, H.",
      "Title": "Machine Learning for Natural Language Processing",
      "PublicationDate": "2022"
    }
  ],
  "Config": {
    "Style": "APA",
    "BilingualAbstract": false,
    "ImagesIndexWithTitle": true
  }
}


*/
