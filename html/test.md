# HTML

[Ref](https://daringfireball.net/projects/markdown/syntax#html)

Note that Markdown formatting syntax is not processed within block-level HTML tags.
E.g., you can’t use Markdown-style *emphasis* inside an HTML block.

Span-level HTML tags — e.g. `<span>`, `<cite>`, or `<del>` — can be used anywhere in a Markdown paragraph,
list item, or header. If you want, you can even use HTML tags instead of Markdown formatting;
e.g. if you’d prefer to use HTML `<a>` or `<img>` tags instead of Markdown’s link or image syntax,
go right ahead.

Unlike block-level HTML tags, Markdown syntax is processed within span-level tags.

## Block level

<div>
This *markup* **should** be treaded as <codeph>HTML</codeph>.
</div>

## Span level

<span>This *markup* **should** be treaded as `markdown`.</span>

<span>
This *markup* **should** be treaded as `markdown`.
</span>

<i>And this **too**.</i>

<i>
And this **too**.
</i>


