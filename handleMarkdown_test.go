package main

import (
    "reflect"
    "testing"
)

func Test_convertLinks(t *testing.T) {


    t.Run("Two_Links", func(t *testing.T) {
markdownExample := `
# Example Markdown

This is a link markdown example:
[Google](https://www.google.com)


Another link:
- [GitHub](https://github.com)

This is an img link:
- ![An Image](https://s3.amazonaws.com/images.seroundtable.com/google-links-1510059186.jpg)

Text without a link here.

`
want:=`
# Example Markdown

This is a link markdown example:
<a href="https://www.google.com">Google</a>


Another link:
- <a href="https://github.com">GitHub</a>

This is an img link:
- <img src="https://s3.amazonaws.com/images.seroundtable.com/google-links-1510059186.jpg" alt="An Image">

Text without a link here.

`
	got := convertLinks(markdownExample)

	if !reflect.DeepEqual(want, got) {
	    t.Errorf("Wrong output: Wanted %s -- Got: %s", want, got)
	}
    })
}

func Test_convertBold(t *testing.T) {
    t.Run("conver_bold", func(t *testing.T){
	sample:=`
This is text with **bold words** and also bold **word**. this is wrong**
*This should mean nothing*.
**
This should also mean nothing**
**
This should also mean nothing
**
`
	want:=`
This is text with <b>bold words</b> and also bold <b>word</b>. this is wrong<b>
*This should mean nothing*.
</b>
This should also mean nothing<b>
</b>
This should also mean nothing
<b>
`
	got:=convertBoldItalic(sample, "**", "<b>")
	
	if !reflect.DeepEqual(want, got) {
	    t.Errorf("\nSample: %s\nWrong output: Wanted %s \nGot: %s",sample, want, got)
	}
    })

    t.Run("conver_bold", func(t *testing.T){
	sample:=`
This is text with *italic words* and also bold <b>word</b>. this is wrong<b>
*This is italic*.
</b>
This should also mean nothing<b>
</b>
This should also mean nothing
<b>
`
	want:=`
This is text with <i>italic words</i> and also bold <b>word</b>. this is wrong<b>
<i>This is italic</i>.
</b>
This should also mean nothing<b>
</b>
This should also mean nothing
<b>
`
	got:=convertBoldItalic(sample, "*", "<i>")
	
	if !reflect.DeepEqual(want, got) {
	    t.Errorf("\nSample: %s\nWrong output: Wanted %s \nGot: %s",sample, want, got)
	}
    })
}
