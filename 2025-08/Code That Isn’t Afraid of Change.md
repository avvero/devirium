![alt text](Код-не-боящийся-изменений4.png)

In this article, you’ll find some personal observations and tips on how to keep a project alive and healthy over the years. No illusions of omniscience. With a touch of healthy cynicism.

My personal experience is in building services as part of product development. In this article, I’m talking specifically about that kind of development - not about creating libraries, frameworks, databases, or other wonderful things. Nor am I touching on project-based development, where code is written to order and then handed off to others for maintenance.

## A Lyrical Prologue

Do you remember when it first appeared? So small, beautiful, not overloaded with features, and untouched by bad advice from Uncle. Do you remember how it timidly rolled out to the staging environment, with the gentle hands of a DevOps helping it take its first steps? And its very first logs? Impossible to understand, but oh, how it kept writing them!

Then you went on vacation. The team promised to look after it, but overfed it with poorly tested features. It no longer fit within its old resource limits, and the containers were bursting at the seams. Years passed. Your sweet little angel turned into a clumsy, gluttonous monster. It stopped cleaning up after itself. It resisted any attempts to teach it new features. Where once a week was enough, now even a month isn’t.

Some people, at that point, leave for another company - hoping to find a good product owner, start new services, and avoid past mistakes. I suggest another path: write code that isn’t afraid of change, preserving its ability to grow without the need for a complete overhaul.

## Ability to Change

Everything discussed in this article can be summed up in one word - [[Evolvability]]. Evolvability is the ability of a system to change: to evolve, adapt, expand, or grow over time without requiring a complete rewrite. What matters is not just whether changes are possible, but how hard those changes are to make. The ideal scenario: adding the same feature should take roughly the same amount of effort in the first month of a project’s life as it does two or three years later. When I talk about code flexibility, I also mean its ability to change.

Evolvability is the opposite of software entropy: as a project ages, its development speed drops, and adding new features or fixing bugs becomes harder and harder. Each step requires more effort and time, which slows development velocity.

## Don’t Plan Too Far Ahead

There’s an old saying: "If you want to make God laugh, tell Him your plans." In software development, it fits perfectly. The most dangerous architectural mistake is trying to design the future in detail when that future is hazy and unknown. We’ve all been there: at the start of a project, trying to account for every possible feature, extension, and load scenario. You think, "What if we need multi-currency support in a year? Better build it in now." Or you decide, "Let’s build for scale from day one - a flexible plugin system, just in case."

A year later, the project has taken an entirely different turn: multi-currency is forgotten, but suddenly you need horoscope integration instead. All that "just in case" code is not only useless - it’s a hindrance. You spend time maintaining code that might be useful someday instead of delivering features that are actually needed.

The YAGNI principle (You Aren’t Gonna Need It) is where true product flexibility shows itself. And if the mention of horoscopes triggered you, that’s exactly the point - you can’t predict the future. Even the wisest developers (yes, even the ones who sneeze right after speaking the truth) can’t tell you where the business will go or what the system will look like in a couple of years. So there’s no point in overcomplicating the design for hypothetical requirements. It’s better to build a minimal working solution today and improve it as real requirements come in.

This isn’t a call to code thoughtlessly - thinking about the future is useful. But there’s a fine line between flexibility and overengineering. You don’t need to create a universal Swiss Army knife right away. You need to make it easy to change later. It’s not about how simple the code is - it’s about how much effort a change will take and what risks it will carry. Code shouldn’t get in the way of future changes; it should support them.

And this especially applies to tests. Jumping ahead a bit, I’ll say: tests play the key role here - more on that later. (Alright, you caught me - that’s one thing I can predict.)

### Don’t Plan Ahead - Even You

A quick side note. Even if you’ve guessed the product’s direction correctly many times and your instincts have become legendary - maybe it’s time to try horse betting. At least there, you’re risking your own money and not someone else’s business. (This is not investment advice.)

## Premature Generality

Product development teaches humility: even the most elegant abstractions quickly run into awkward edge cases. More often than not, plain duplication is better than premature generalization.

A classic example: you get a request -  "We need another slightly different report format." You already have a `ReportGenerator` class with a couple of subclasses for different formats. The temptation is strong - create a third subclass, or better yet, design a universal mega-class with plugins for any format. But if these reports truly differ by just a couple of fields, it’s often faster and more reliable to… yes, copy the existing generator code and tweak it for the new case. Sure, it sounds like "embracing copy-paste." But you’ll get the correct result quickly, without breaking existing logic.

Anticipating the cries about maintainability: yes, code duplication has a cost. Two similar modules will need to be updated in parallel for shared changes. But in my experience, that’s the lesser evil compared to premature abstraction. An overly clever, generalized system often turns out inflexible. The moment you encounter a case that doesn’t fit your elegant hierarchy, you’ll spend more time reshaping your abstractions than you would have spent maintaining two independent, simple solutions.

The golden rule here is simple: abstract only after something has happened three times. In other words:
- Once  -  no reason to abstract at all.
- Twice  -  a weak signal; a third time may be coming.
- Thrice  -  it’s time to extract a common mechanism.

Following this rule helps avoid the bad habit of "writing a framework for a single task." I’ve seen it described in various sources under names like The Rule of Three Envelopes, Triangulation, 1-2-refactor, or simply The Rule of Three.

## Accept It: You Won’t Understand Everything Right Away

One of the key truths: you won’t know everything from the start. Big features often begin with vague requirements. Half of them are incomplete, and the other half only seem clear. As a result, the first version of the system is built almost blindly - and that’s fine. The goal isn’t to get it right the first time, but to make it possible to redo it later.

Many developers fear rewriting "what’s already done." The time spent feels too precious to throw away. But often, discarding and rewriting is faster and safer. Code isn’t carved in stone, even if we strive to write it beautifully and cleanly. If it turns out the current solution doesn’t scale well or fails to cover an important new scenario, it’s better to admit it and rework it than to heroically drag along a flawed implementation.

In practice, code flexibility means that refactoring and redesign are routine, not disasters. Your code should be ready for it. The architecture should allow components, modules, or algorithms to be swapped out with minimal effort. If a module is designed so that removing it requires rewriting the entire system - that’s a red flag.

Accepting the limits of your knowledge leads to more adaptable design: you don’t build for specific anticipated extensions, you build for change in general. That means shifting from "we accounted for all possible scenarios in advance" to "we made it so that adding a new scenario is relatively painless."

## Iterations Over Revolutions

A product’s life is a continuous sequence of iterations. Instead of shipping a massive release every six months, teams tweak screws weekly - or even daily - adding small features, fixing bugs, and making minor improvements. This pace shapes the development approach: decisions are quick and localized. Code is written to "solve the problem now," with the understanding that another developer may come along later and rewrite it to suit new needs.

Speaking of developers - over the product’s lifetime, you’ll probably see a whole parade of them. Each will bring their own style and ideas. Somewhere deep in the project are files whose authors barely remember what they wrote back in 2017. In such conditions, trying to "plan everything in advance" is especially pointless: a newcomer with fresh eyes will look at your carefully crafted architecture and think, "Why is this so complicated? Let’s simplify." Then they’ll rework half the system.

Iterativeness means being ready to change course in small steps. It’s like adjusting a ship’s heading: small maneuvers instead of a sharp 90-degree turn. A project that evolves iteratively is easier to adapt - each small improvement tests hypotheses and allows for timely course correction. In contrast, trying to lock in a finished architecture from the start often ends in massive refactoring - or even a rewrite - once it becomes clear reality doesn’t match the plan.

## Feature Grafting

In gardening, there’s a technique called "grafting": attaching a branch from one plant onto the trunk of another. Software development sometimes works the same way. When you want to quickly grow a new feature, it’s often easier to graft it into an existing service rather than immediately building a separate module or microservice. You embed the new functionality "on foreign territory," taking advantage of the resources and context of existing code. Once that branch grows strong, takes root, and proves its value, you can carefully separate it and transplant it into its own "pot" (a dedicated service or module).

This approach helps avoid premature complexity. It’s tempting to fear touching an old monolith or large service, and to push every new idea into a separate place "so as not to break anything." But the truth is that early-stage internal integration is often faster and more reliable. You already have the infrastructure, configurations, and database - why create extra entities if you can quickly plug into what exists?

Of course, this approach has downsides: the codebase temporarily swells, boundaries of responsibility blur, and old code gets "extra baggage." But if you keep the end goal in mind (extracting the feature into its own component), that debt can be paid off gradually. In return, you validate the idea with minimal cost, and later refactor cleanly, removing the grafted code from its original host.

## API Is Not Set in Stone

Ever notice how some internal API docs say, "This method is deprecated, use the new one" - and keep saying that for several versions in a row? That’s just evolution at work. No matter how much you try to design interfaces and contracts perfectly, reality proves one thing: APIs change. Maybe less often than internal logic, but they still change.

This is especially true for internal APIs between services. There’s no such thing as going five years without a single endpoint signature changing. Even if the contract itself stays stable, new versions appear, old ones get phased out, or new parameters get added. A mature product lives in a state of constant restructuring, and APIs are part of that life.

So how do you live with it? First - versioning. Accept that sooner or later you’ll have to roll out v2, v3… It’s better to design for multiple API versions upfront, so you’re not breaking clients every time. Second - deprecation policy. Have a strategy for retiring old functionality, even if it’s just marking it "Deprecated" and giving it a year of support before removal. That’s far better than trying to keep the first version alive forever out of fear of user backlash. And users adapt - better to give them a clear plan ("The old API will be shut down in N months, please migrate to the new one") than to silently change behavior and then awkwardly avoid the topic.

One more thing - flexibility in integrations. If your product interacts with external APIs, be ready to rewrite things when the partner updates theirs. Budget time for it in your planning. Few things stall evolution as effectively as a suddenly broken integration you didn’t allocate resources to fix.

## Tests: Safety and Freedom

I’ll never get tired of saying it: tests are a flexible code’s best friends. When you have hundreds of modules and services, it’s all too easy to accidentally break a contract or a dependency. Tests need maintenance too, and bad tests can hinder progress as much as bad code. But living without at least some tests is far scarier. Without them, every refactoring becomes a game of roulette: "Will it break or won’t it?" And fear of changing code is the number one killer of flexibility.

My personal choice - proven over years of working on various services - is the Testing Trophy approach. It focuses on integration tests centered around observable behavior. When tests are decoupled from the implementation, you can change that implementation every day and still remain safe.

I also want to call out TDD (Test-Driven Development). Some debate its overall value as a methodology, but for me there’s one major benefit: it separates the effort of writing code to fulfill a contract from the effort of shaping the code’s design. You’re not trying to nail both the perfect design and the correct behavior in one pass - you split those tasks. This gives your code room to rest and mature while awaiting refactoring. In the context of our topic, that means fewer chances to overbuild or overlook something important.

### One Pile

That bit about "letting the code rest" wasn’t just a throwaway line. There’s an interesting practice called One Pile, described in Kent Beck’s book Tidy First?. The idea is that understanding code is harder than writing it. If the code is split into too many tiny pieces, it can be useful to first combine it into one whole to see the overall structure, and only then break it back down into methods and modules. This is especially helpful in the early stages, when the task and its boundaries are still unclear.

## Conclusion

The ability to change isn’t about a specific language or framework - it’s about an approach to development. It’s about building systems that can shed their skin, drop their tail, grow new limbs - and still keep doing their job. There’s no silver bullet, but there is a mindset: don’t carve everything in granite, be ready for change, stay open to it, and don’t fall so in love with your own solutions that you become blind to better ones.

#article #ignore