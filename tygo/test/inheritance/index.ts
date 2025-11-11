export interface Base {
  name: string;
}
export interface Base2 {
  id: any;
}
export interface Base3 {
  class: string;
  level: number /* int */;
}
export interface Other extends Base, Base2, Partial<Base3>, bookapp.Book, bookapp.TextBook {
  otherWithBase: Base;
  otherWithBase2: Base2;
  otherValue: string;
  author: bookapp.AuthorWithInheritance;
}
