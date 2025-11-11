package inheritance

import (
	bookapp "github.com/codeliger/tygo/examples/bookstore"
)

type Other struct {
	*Base          `tstype:",extends,required"`
	Base2          `tstype:",extends"`
	*Base3         `tstype:",extends"`
	OtherWithBase  Base                          `                                          json:"otherWithBase"`
	OtherWithBase2 Base2                         `                                          json:"otherWithBase2"`
	OtherValue     string                        `                                          json:"otherValue"`
	Author         bookapp.AuthorWithInheritance `tstype:"bookapp.AuthorWithInheritance" json:"author"`
	bookapp.Book   `tstype:",extends"`
	TextBook       *bookapp.TextBook `tstype:",extends,required"`
}
