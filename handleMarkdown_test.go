package main

import (
	"fmt"
	"reflect"
	"testing"
)

func checkDiff(want, got, sample string) string {
	if !reflect.DeepEqual(want, got) {
		diff := markDiffSpace(want, got)
		if diff != "" {
			return fmt.Sprintf(
				"\nSample: \n%s\nWanted: \n@%s@ \n   Got: \n%s\nDiff:\n%s\n",
				sample,
				want,
				got,
				diff,
			)
		}
		return fmt.Sprintf(
			"\nSample: \n%s\nWanted: \n@%s@ \n   Got: \n@%s@",
			sample,
			want,
			got,
		)
	}
	return ""
}

func markDiffSpace(w, g string) (diff string) {
	if len(w) != len(g) || len(w) < 1 || len(g) < 1 {
		return ""
	}
	result := []rune(g)
	s := byte(' ')
	for idx := range w {
		if w[idx] == s && g[idx] != s {
			result[idx] = '@'
		}
	}
	if string(result) == g {
		return ""
	}
	return string(result)
}

func Test_linksOrImgToHtml(t *testing.T) {
	t.Run("links_to_html", func(t *testing.T) {
		sample := `
# Example Markdown

This is a link markdown example:
[Google](https://www.google.com)


Another link:
- [GitHub](https://github.com)

This is an img link:
- ![An Image](https://s3.amazonaws.com/images.seroundtable.com/google-links-1510059186.jpg)

Text without a link here.

`
		want := `
# Example Markdown

This is a link markdown example:
<a href="https://www.google.com">Google</a>


Another link:
- <a href="https://github.com">GitHub</a>

This is an img link:
- ![An Image](https://s3.amazonaws.com/images.seroundtable.com/google-links-1510059186.jpg)

Text without a link here.

`
		got := linkToHtml(sample)

		diff := checkDiff(want, got, sample)
		if diff != "" {
			t.Error(diff)
		}
	})
	t.Run("img_to_html", func(t *testing.T) {
		sample := `
# Example Markdown

This is a link markdown example:
[Google](https://www.google.com)


Another link:
- [GitHub](https://github.com)

This is an img link:
- ![An Image](https://s3.amazonaws.com/images.seroundtable.com/google-links-1510059186.jpg)

Text without a link here.

`
		want := `
# Example Markdown

This is a link markdown example:
[Google](https://www.google.com)


Another link:
- [GitHub](https://github.com)

This is an img link:
- <img src="https://s3.amazonaws.com/images.seroundtable.com/google-links-1510059186.jpg" alt="An Image">

Text without a link here.

`
		got := imgToHtml(sample)

		diff := checkDiff(want, got, sample)
		if diff != "" {
			t.Error(diff)
		}
	})
}

func Test_boldToHtml(t *testing.T) {
	t.Run("conver_bold", func(t *testing.T) {
		sample := `
This is text with **bold words** and also bold **word**. this is wrong**
*This should mean nothing*.
**
This should also mean nothing**
**
This should also mean nothing
**
`
		want := `
This is text with <b>bold words</b> and also bold <b>word</b>. this is wrong**
*This should mean nothing*.
**
This should also mean nothing**
**
This should also mean nothing
**
`
		got := boldToHtml(sample)

		diff := checkDiff(want, got, sample)
		if diff != "" {
			t.Error(diff)
		}
	})

	t.Run("conver_italic", func(t *testing.T) {
		sample := `
This is text with *italic words* and also bold <b>word</b>. this is wrong<b>
*This is italic*.
</b>
This should also mean nothing<b>
</b>
This should also mean nothing
<b>
`
		want := `
This is text with <i>italic words</i> and also bold <b>word</b>. this is wrong<b>
<i>This is italic</i>.
</b>
This should also mean nothing<b>
</b>
This should also mean nothing
<b>
`
		got := italicToHtml(sample)

		diff := checkDiff(want, got, sample)
		if diff != "" {
			t.Error(diff)
		}
	})
}

func Test_allItalicAndBoldToHtml(t *testing.T) {
	t.Run("Normal string", func(t *testing.T) {
		sample := `***bold text***`
		want := `<em><strong>bold text</strong></em>`
		got := allItalicAndBoldToHtml(sample)
		diff := checkDiff(want, got, sample)
		if diff != "" {
			t.Error(diff)
		}
	})
	t.Run("same line", func(t *testing.T) {
		sample := `***bold text*** ***more bold text***`
		want := `<em><strong>bold text</strong></em> <em><strong>more bold text</strong></em>`
		got := allItalicAndBoldToHtml(sample)
		diff := checkDiff(want, got, sample)
		if diff != "" {
			t.Error(diff)
		}
	})
	t.Run("first space", func(t *testing.T) {
		sample := `*** bold text***`
		want := `*** bold text***`
		got := allItalicAndBoldToHtml(sample)
		diff := checkDiff(want, got, sample)
		if diff != "" {
			t.Error(diff)
		}
	})
	t.Run("second spcae", func(t *testing.T) {
		sample := `***bold text ***`
		want := `***bold text ***`
		got := allItalicAndBoldToHtml(sample)
		diff := checkDiff(want, got, sample)
		if diff != "" {
			t.Error(diff)
		}
	})
	t.Run("line break", func(t *testing.T) {
		sample := `
***
***`
		want := `
***
***`
		got := allItalicAndBoldToHtml(sample)
		diff := checkDiff(want, got, sample)
		if diff != "" {
			t.Error(diff)
		}
	})
}

func Test_oListToHtml(t *testing.T) {
	t.Run("Two ordered lists", func(t *testing.T) {
		sample := `
2. This should be a 2
5. This should be a 3
6. This should be a 4
List interrupted
20. This should be a 20
25. This should be a 21
`
		want := `
<ol>
    <li value="2">This should be a 2</li>
    <li>This should be a 3</li>
    <li>This should be a 4</li>
</ol>
List interrupted
<ol>
    <li value="20">This should be a 20</li>
    <li>This should be a 21</li>
</ol>
`
		got := oListToHtml(sample)
		diff := checkDiff(want, got, sample)
		if diff != "" {
			t.Error(diff)
		}
	})
}

func Test_uListToHtml(t *testing.T) {
	t.Run("simple list", func(t *testing.T) {
		sample := `- <a href="google.com">Google</a>`
		want := `<ul>
    <li> <a href="google.com">Google</a></li>
</ul>`
		got := uListToHtml(sample)
		diff := checkDiff(want, got, sample)
		if diff != "" {
			t.Error(diff)
		}
	})
}
