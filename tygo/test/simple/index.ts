export type MyUint8 = number /* uint8 */;
export type MyInt = number /* int */;
export type MyString = string;
export type MyAny = any;
/**
 * Should be a number in TypeScript.
 */
export type MyRune = number /* rune */;
/**
 * Comment for a struct
 */
export interface MyStruct {
  some_field: any;
  /**
   * Comment for a field
   */
  OtherField: boolean; // Comment after line
  FieldWithImportedType: any /* some.Type */;
}
/**
 * Foo
 */
export type MyValue = number /* int */; // Bar
/**
 * Mixed unexported and exported
 */
export const ExportedValue = 42; // A comment on an exported constant
export interface A {
  Foo: string; // A comment on fields separated by a comma
  Bar: string; // A comment on fields separated by a comma
}
export interface B {
  /**
   * A comment above the fields separated by a comma
   */
  Foo: string;
  /**
   * A comment above the fields separated by a comma
   */
  Bar: string;
}
export const Pi = 3.14; // A comment on constants separated by a comma
export const E = 2.71; // A comment on constants separated by a comma
