![alt text](DigitalGarden.jpg)

The topic of note-taking remains relevant today. We know the benefits it provides to the author. We are familiar with various approaches to note-taking, tools that can be used, and we have choices. Imagine you have found your approach, your tool, and your note base is growing and pleasing to the eye. What next? I want to discuss one path of development in this area.

This article is dedicated to the concept of a digital garden — the philosophy of publicly maintaining personal notes.

## My Path

For the past 20 years, I have kept notes using methods that lacked systematization, reliability, or usefulness: paper notebooks, text files, Evernote, and other applications whose names have faded from memory. Three years ago, I started using the Zettelkasten style (as I understood and adapted it for myself). I learned about this approach from [Rob Muhlestein](https://github.com/rwxrob) whom I encountered on Twitch. Experiments led me to my current note-taking method: in the form of Markdown files stored in a Git repository. On my computer, I work with them in VSCode, and on my phone, I use the default "Notes" app. Why not Obsidian? I tried it and decided to stick with VSCode for the following reasons:
1. VSCode is a basic editor that is always open and used for work files.
2. I could not set up Obsidian synchronization via Git, and iCloud synchronization caused app crashes on my phone.
3. I need a quickly accessible app on my phone to jot down fleeting thoughts before, for example, reaching for a roll of toilet paper and experiencing discomfort from waiting. At the same time, I need the note base on my phone much less frequently than the need to urgently write something down.

Separating apps for sudden thoughts and the main base has an advantage: it creates a ritual of transferring notes from the phone app to the main base. I review fresh notes, tag them, and add details. This allows me to revisit the recorded thought and increase the chances of not forgetting it.

## What is a Digital Garden

The phrase "digital garden" is a metaphor describing an approach to note-taking. It is not just a set of tools like WordPress plugins or Jekyll templates. The idea of a garden is familiar to all of us — it is a place where something grows. Gardens can be very personal and filled with gnome figurines, or they can be sources of food and vitality. And who knows what a sudden visitor to your garden might see? You in a lovely pajama with a glass of fresh juice under an apple tree? Or perhaps standing upside down, trying to bring a bit of order and pull out weeds?

The digital gardening metaphor emphasizes the slow growth of ideas through writing, rewriting, editing, and revisiting thoughts in a public space. Instead of fixed opinions that never change, this approach allows ideas to develop over time. The goal of digital gardening is to use the collective intelligence of your network to create constructive feedback loops. If done right, you will have an accessible representation of your thoughts that can be "sent out" into the world, and people will be able to respond to it. Even for the most raw ideas, it helps to create a feedback loop to strengthen and fully develop the idea.

Core principles of gardening:
1. Connections over timelines. Gardens are organized around contextual and associative connections; concepts and themes within each note define how it relates to others. The publication date is not the most important aspect of the text.
2. Continuous growth. Gardens never end; they are constantly growing, evolving, and changing, like a real garden.
3. Imperfection. Gardens are inherently imperfect. They do not hide their rough edges and do not claim to be a permanent source of truth.
4. Learning in public. To create constructive feedback loops.
5. Personal and experimental. Gardens are inherently heterogeneous. You may plant the same seeds as your neighbor but get a different arrangement of plants. You organize the garden around ideas and means that fit your way of thinking rather than a standard template.
6. Independent ownership. Gardening is about creating your own little corner of the internet that you fully control.

### Learning in Public

This point may raise questions, as it suggests sharing something for free, i.e., as a gift. The approach implies that you publicly document your steps, thoughts, mistakes, and successes in mastering a new topic or skill. This allows you not only to share the result but also to demonstrate the thinking and learning process, which can be useful to others.

Besides altruistic motives, personal interest is worth noting. Sometimes I am lazy to clearly formulate thoughts in notes, which then backfires — I can't understand what I meant. It's amusing how I always rely on my future superpower to decipher my own nonsensical notes — a belief that remains unshakable, though entirely unfounded. But when it comes to others, I don’t allow such naivety. Knowing that my note might not only attract attention but also genuinely help someone motivates me to make an effort and articulate the thought properly.

## How to Share

Today, convenient tools have made creating a fully customizable website much easier. Services like Netlify and Vercel have removed deployment complexities. Static site generators like Jekyll, Gatsby, 11ty, and Next simplify creating complex sites that automatically generate pages and handle load time, image optimization, and SEO.

Obsidian offers the ability to publish notes through its subscription platform. Using this service does not feel like "independent ownership."

I chose [Quartz](https://quartz.jzhao.xyz/) for myself. It is a free static generator based on Markdown content. Quartz is designed primarily as a tool for publishing digital gardens on the internet. It is simple enough for people without technical experience but powerful enough for customization by experienced developers.

## My Digital Garden

As mentioned, I store my notes in a public Git repository. Changes are automatically published on GitHub Pages and available on [my digital garden page](https://devirium.avvero.pw). The design, solution scheme, and GitHub Actions scripts are available for review in [this note](https://devirium.avvero.pw/2024/2024-07/How-I-Built-Devirium/) if you want to create something similar. Some use RSS for updates, while I use a Telegram channel for this purpose. It is specifically an update channel, messages are not posted separately.

## Materials for Further Study

The concept has a philosophy and history; I won't retell them but will provide links where you can read more: [https://maggieappleton.com](https://maggieappleton.com). Examples of digital gardens can be found at [https://github.com/MaggieAppleton/digital-gardeners](https://github.com/MaggieAppleton/digital-gardeners) and [https://github.com/jackyzha0/quartz/blob/v4/docs/showcase.md](https://github.com/jackyzha0/quartz/blob/v4/docs/showcase.md), as well as through GitHub search if the repository is tagged with the relevant topic — [https://github.com/topics/digital-garden](https://github.com/topics/digital-garden).

## Conclusion

Establishing a digital garden has been a logical continuation of my journey, started with the Zettelkasten method. How has it affected me? After the initial effort to set up and deploy the system, I hardly maintain it except for occasional issues. And now I continue to push my notes to Git. The only thing is, I have started to make them more understandable.

Thank you for reading the article, and good luck in your quest for organizing thoughts and creating an effective space for ideas!