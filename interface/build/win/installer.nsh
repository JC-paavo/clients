!macro customHeader
  RequestExecutionLevel admin
!macroend

!macro customInstall
    WriteRegStr HKCR "jms" "" "URL:jms"
    WriteRegStr HKCR "jms" "URL Protocol" ""
    WriteRegStr HKCR "jms\shell" "" ""
    WriteRegStr HKCR "jms\shell\open" "" ""
    WriteRegStr HKCR "jms\shell\open\command" "" '$INSTDIR\resources\bin\windows\JumpServerClient.exe "%1"'
!macroend

!macro customUnInstall
    DeleteRegKey HKCR "jms"
!macroend