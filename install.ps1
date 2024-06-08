# Download the latest version of the TU executable

Invoke-WebRequest -Uri "https://github.com/mistweaverco/tu/releases/latest/download/tu.exe" -OutFile "tu.exe"

# Move the TU executable to the System32 directory

Move-Item -Path "tu.exe" -Destination "C:\Windows\System32\tu.exe"
