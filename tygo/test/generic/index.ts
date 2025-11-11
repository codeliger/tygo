/**
 * Comment for UnionType
 */
export type UnionType = unknown;
export type Derived = unknown;
export type Any = unknown;
export type Empty = unknown;
export type Something = unknown;
export interface EmptyStruct {
}
export interface ValAndPtr {
  Val: unknown;
  /**
   * Comment for ptr field
   */
  Ptr?: unknown; // ptr line comment
}
export interface ABCD {
  a: string;
  b: string;
  c: UnionType;
  d: unknown;
}
export interface Foo {
  Bar: unknown;
  Boo?: unknown;
}
export interface WithFooGenericTypeArg {
  some_field: Foo;
}
export interface Single {
  Field: unknown;
}
export type SingleSpecific = Single;
export interface AnyStructField {
  Value: unknown;
  SomeField: string;
}
export type JsonArray = unknown[];
