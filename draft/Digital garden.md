Тема ведения заметок не теряет свою актуальность и по сей день. У нас есть представление о том, что это дает автору. Нам знакомы различные подходы к ведению заметок. Какие инструменты для этого использовать мы уже тоже знаем и можем выбирать. Но вот представьте, что вы нашли свой подход, нашли свой инструмент и база заметок растет и радует глаз. А что дальше? И вот об одном из путей развития темы заметок я бы хотел рассказать.

В данной статье пойдет речь о концепции Цифрового сада - философии публичного ведения базы личных заметок.

## Мой путь

Последние 20 лет я вел заметки способами, не отличающимися ни систематичностью ни надежностью ни пользой - бумажные записные книжки, txt файлы, Evernote и прочите приложения, названия которых уже выветрились из памяти. И вот уже как три года я веду заметки в стиле Zettelkasten (на сколько я его понял и принял к использованию с поправками для себя). Узнал я о таком подходе к ведению от Rob Muhlestein - https://github.com/rwxrob, на которого наткнулся на твиче. Последующие после знакомства эксперименты привели меня к сегодняшнему способу. Заметки я веду в виде md файлов и храню в git репозитории. На компьютере работаю с ними в VSCode, на телефоне использую стандартное приложение "Заметки". Почему не Obsidian, спросите вы? Я пробовал и решил остаться на VSCode по следующим причинам:
1. VSCode это базовый редактор, который у меня итак всегда открыт и постоянно используется, в том числе для ведения рабочих файлов.
2. Не смог подружить Obsidian с синхронизацией через git, а синхронизация через iCloud вызывает зависания в приложении на телефоне. 
3. Мне необходимо максимально быстро открывающееся приложение на телефоне, чтобы я смог в любой момент успеть и записать ускользающую мысль, перед тем как например потянуться к рулону туалетной бумаги и не испытывать дискомфорт от ожидания. При этом наличие базы заметок в телефоне мне требуется значительно реже, чем необходимость что-то срочно записать. 

К слову, в разделении приложений для записи внезапных мыслей и основной базы есть один плюс - появляется ритуал переноса заметок из приложения на телефоне в основную базу, я пересматриваю свежие заметки, проставляю теги, дополняю деталями. Это позволяет лишний раз соприкоснуться с записанной мыслью и увеличить шанс ее не забыть.

## Что такое цифровой сад

Фраза "цифровой сад" – это метафора, которая описывает подход к ведению заметок, где акцент делается не на конечный "отполированный" результат, а на процесс. Идея сада, я думаю нам всем знакома - это обычно место, где что-то растет. Сады могут быть очень личными и наполненными фигурками гномов из гипса, или они могут быть источником пищи и жизненной силы. Метафора мне кажется мне очень удачной, кто знает, что увидит внезапный гость вашего сада? Вас в прекрасной пижаме с бокалом свежего сока под яблоней? Или быть может стоящим кверху задом, пытающегося навести хоть немного порядка и вырвать сорняки?

Цифровой сад это не набор специфичных инструментов, типа плагинов для Wordpress или шаблонов Jekyll. 


Это отличает сад от персонального блока, как мы обычно его пониманием, который содержит набор отполированных и выверенных статей.

>My small collection highlighted a number of sites that are taking a new approach to the way we publish personal knowledge on the web.
>They're not following the conventions of the "personal blog," as we've come to know it. Rather than presenting a set of polished articles, displayed in reverse chronological order, these sites act more like free form, work-in-progress wikis.
>Caufield makes clear digital gardening is not about specific tools – it's not a Wordpress plugin, Gastby theme, or Jekyll template. It's a different way of thinking about our online behaviour around information - one that accumulates personal knowledge over time in an explorable space.
>Gardens present information in a richly linked landscape that grows slowly over time.
>Everything is arranged and connected in ways that allow you to explore.
>Joel focused on the process of digital gardening, emphasising the slow growth of ideas through writing, rewriting, editing, and revising thoughts in public. Instead of slapping Fully Formed Opinions up on the web and never changing them.

The goal of digital gardening should be to tap into your network’s collective intelligence to create constructive feedback loops. If done well, I have a shareable representation of my thoughts that I can send out into the world and people can respond. Even for my most half-baked thoughts, this helps me create a feedback cycle to strengthen and fully flesh out that idea.

>Не одноразовые
>Вeчноцветущие 

> The phrase "digital garden" is a metaphor for thinking about writing and creating that focuses less on the resulting "showpiece" and more on the process, care, and craft it takes to get there. While not everybody has or works in a dirt garden, we all share a familiarity with the idea of what a garden is.
>A garden is usually a place where things grow.
>Gardens can be very personal and full of whimsy or a garden can be a source of food and substance.

https://www.swyx.io/learn-in-public


The Six Patterns of Gardening
1. Topography over Timelines. Gardens are organised around contextual relationships and associative links; the concepts and themes within each note determine how it's connected to others. Gardens don't consider publication dates the most important detail of a piece of writing. Dates might be included on posts, but they aren't the structural basis of how you navigate around the garden.
2. Continuous Growth. Gardens are never finished, they're constantly growing, evolving, and changing. Just like a real soil, carrot, and cabbage garden.
3. Imperfection & Learning in Public. Gardens are imperfect by design. They don't hide their rough edges or claim to be a permanent source of truth.
4. Playful, Personal, and Experimental. Gardens are non-homogenous by nature. You can plant the same seeds as your neighbour, but you'll always end up with a different arrangement of plants. You organise the garden around the ideas and mediums that match your way of thinking, rather than off someone else's standardised template.
5. Intercropping & Content Diversity
6. Independent Ownership. Gardening is about claiming a small patch of the web for yourself, one you fully own and control.

"Digital Garden" is a philosophy for sharing personal knowledge on the web:
- Gardens are organized around contextual relationships and associative links.
- Gardens are never finished; they're constantly growing, evolving, and changing.
- Gardens are imperfect by design. They don't hide their rough edges or claim to be a permanent source of truth.
- Gardens are meant to be shared.

## Зачем делиться / Learning in Public

Open Source Your Knowledge

Зачем это мне. 
1. Мне нравится как другие ведут сады, мне интересно их читать, я хотел бы делать так же.
2. Publishing imperfect and early ideas requires that we make the status of our notes clear to readers.
3. A habit of creating learning exhaust. Whatever your thing is, make the thing you wish you had found when you were learning. https://www.swyx.io/learn-in-public
4. Получить фидбек - The goal of digital gardening should be to tap into your network’s collective intelligence to create constructive feedback loops.

## Как делиться

The current state of web development helped here too. While it feels like we've been in a slow descent into a horrifyingly complex and bloated web development process, a number of recent tools have made it easier to get a fully customised website up and running. Services like Netlify and Vercel have taken the pain out of deployment. Static site generators like Jekyll, Gatsby, 11ty and Next make it easier to build sophisticated websites that auto-generate pages, and take care of grunt work like optimising load time, images, and SEO.

These services are trying to find a happy middle ground between tediously hand-coding solutions, and being trapped in the restrictions of Wordpress or Squarespace.

Tools like Obsidian, TiddlyWiki, and Notion are all great options. Many of them offer fancy features like nested folders, Bi-Directional Links, footnotes, and visual graphs.

However, many of these no-code tools still feel like cookie-cutter solutions. Rather than allowing people to design the information architecture and spatial layouts of their gardens, they inevitably force people into pre-made arrangements. This doesn't meant they don't "count,” as "real” gardens, but simply that they limit their gardeners to some extent. You can't design different types of links, novel features, experimental layouts, or custom architecture. They're pre-fab houses instead of raw building materials.

Для себя я выбрал кварц.

Quartz is designed first and foremost as a tool for publishing digital gardens to the web. To me, digital gardening is not just passive knowledge collection. It’s a form of expression and sharing.

## Мой цифровой сад

Недавно я открыл для себя концепцию цифрового сада — философию делиться личными знаниями в интернете. Сады постоянно растут, изменяются и предназначены для общего доступа. Подробнее можно почитать здесь https://maggieappleton.com/garden-history. Вдохновившись этой концепцией, я создал свой цифровой сад — https://devirium.avvero.pw и опубликовал все свои заметки. Также создал канал в Телеграме: https://t.me/devirium, куда синхронизируются все заметки. Использую его как средство доступа — Телеграм часто открыт, он очень быстрый и есть поиск по тегам. Раз в день туда прилетает случайная заметка, которую я читаю и вспоминаю. На главной странице сада приведена дизайн схема и все компоненты, если вы захотите сделать себе подобное.

## References and Further Reading

У концепции есть философия и история, я не буду их пересказывать, лишь представлю ссылки на то, где про это можно почитать подбронее - https://maggieappleton.com.

From [Philosophy of Quartz](https://quartz.jzhao.xyz/philosophy)

## Цифровые сады известных людей

Известные мне:
- https://www.dschapman.com/notes/
- [[Yenly]]
- https://jzhao.xyz
- https://github.com/MaggieAppleton/digital-gardeners
- https://digitalgarden.guidopercu.dev/
- https://joelhooks.com/digital-garden/
- https://oliverfalvai.com/evergreen/my-quartz-+-obsidian-note-publishing-setup
- https://oliverfalvai.com/Evergreen/Inspiring-gardens

- [Quartz Documentation (this site!)](https://quartz.jzhao.xyz/)
- [Jacky Zhao's Garden](https://jzhao.xyz/)
- [Socratica Toolbox](https://toolbox.socratica.info/)
- [oldwinter の数字花园](https://garden.oldwinter.top/)
- [Aaron Pham's Garden](https://aarnphm.xyz/)
- [The Quantum Garden](https://quantumgardener.blog/)
- [Abhijeet's Math Wiki](https://abhmul.github.io/quartz/Math-Wiki/)
- [Matt Dunn's Second Brain](https://mattdunn.info/)
- [Pelayo Arbues' Notes](https://pelayoarbues.github.io/)
- [Vince Imbat's Talahardin](https://vinceimbat.com/)
- [🧠🌳 Chad's Mind Garden](https://www.chadly.net/)
- [Pedro MC Fernandes's Topo da Mente](https://www.pmcf.xyz/topo-da-mente/)
- [Mau Camargo's Notkesto](https://notes.camargomau.com/)
- [Caicai's Novels](https://imoko.cc/blog/caicai/)
- [🌊 Collapsed Wave](https://collapsedwave.com/)
- [Sideny's 3D Artist's Handbook](https://sidney-eliot.github.io/3d-artists-handbook/)
- [Mike's AI Garden 🤖🪴](https://mwalton.me/)
- [Brandon Boswell's Garden](https://brandonkboswell.com)
- [Scaling Synthesis - A hypertext research notebook](https://scalingsynthesis.com/)
- [Data Dictionary 🧠](https://glossary.airbyte.com/)
- [sspaeti.com's Second Brain](https://brain.sspaeti.com/)
- [🪴Aster's notebook](https://notes.asterhu.com)
- [🥷🏻🌳🍃 Computer Science & Thinkering Garden](https://notes.yxy.ninja)
- [A Pattern Language - Christopher Alexander (Architecture)](https://patternlanguage.cc/)
- [Gatekeeper Wiki](https://www.gatekeeper.wiki)

If you want to see your own on here, submit a [Pull Request adding yourself to this file](https://github.com/jackyzha0/quartz/blob/v4/docs/showcase.md)!

## Заключение 

Теги: obsidian, организация работы, заметки

#article #draft #garden