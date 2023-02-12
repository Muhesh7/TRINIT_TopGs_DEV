from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class AppAuthRequest(_message.Message):
    __slots__ = ["app_id", "app_secret"]
    APP_ID_FIELD_NUMBER: _ClassVar[int]
    APP_SECRET_FIELD_NUMBER: _ClassVar[int]
    app_id: str
    app_secret: str
    def __init__(self, app_id: _Optional[str] = ..., app_secret: _Optional[str] = ...) -> None: ...

class AppAuthResponse(_message.Message):
    __slots__ = ["isSuccess", "rule"]
    ISSUCCESS_FIELD_NUMBER: _ClassVar[int]
    RULE_FIELD_NUMBER: _ClassVar[int]
    isSuccess: bool
    rule: _containers.RepeatedCompositeFieldContainer[Rule]
    def __init__(self, isSuccess: bool = ..., rule: _Optional[_Iterable[_Union[Rule, _Mapping]]] = ...) -> None: ...

class Rule(_message.Message):
    __slots__ = ["match_type", "parameter"]
    MATCH_TYPE_FIELD_NUMBER: _ClassVar[int]
    PARAMETER_FIELD_NUMBER: _ClassVar[int]
    match_type: str
    parameter: str
    def __init__(self, parameter: _Optional[str] = ..., match_type: _Optional[str] = ...) -> None: ...
