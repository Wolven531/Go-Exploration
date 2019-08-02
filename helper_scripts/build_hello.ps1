# NOTE: MUST be run from working directory `helper_scripts`

Set-Location -Path '../src/hello/'
Invoke-Expression -Command 'go build'

Write-Host -ForegroundColor Green 'You can now run `hello.exe` from this location'
Write-Host -ForegroundColor Yellow '(or `cd ..\..\helper_scripts` to return)'
