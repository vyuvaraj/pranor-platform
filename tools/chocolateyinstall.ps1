$packageName = 'pranor'
$fileType = 'exe'
$silentArgs = '/VERYSILENT /SUPPRESSMSGBOXES /NORESTART'
$url = 'https://github.com/vyuvaraj/pranor-repo/releases/download/v1.7.0/Pranor-windows-setup.exe'

Install-ChocolateyPackage -PackageName $packageName `
                          -FileType $fileType `
                          -SilentArgs $silentArgs `
                          -Url $url
