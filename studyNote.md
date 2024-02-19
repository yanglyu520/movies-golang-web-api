## learnings:
1. Question: value receiver or pointer receiver for a function?
   Answer: The rule about pointers vs. values for receivers is that value methods can be invoked on pointers and values, but pointer methods can only be invoked on pointers.

2. Question: What do you need to know when using json.Newdecorder(r.Body).Decode(&input)?
Answer:
-  It is hard to know if ciient deliberately send a 0 value in the int, or if we omit a particular key/value pair when providing the request body
-  when calling Decode(), you must pass a non-nil ponter as the target decode destination. If you dont use a pointer, it will return a json.InvalidUnmarshalError error at runtime
- go's json.Decoder provides a DisallowUnkownFields() setting that we can use to generate an error when unknow fields are included in the input