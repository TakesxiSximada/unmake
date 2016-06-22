mkdir -p unmake
curl -L -o unmake/Makefile https://raw.githubusercontent.com/TakesxiSximada/unmake/develop/unmake/Makefile
echo "Insert next lines in your Makefile top"
echo "#-------------------------------------"
echo "UNMAKE_CURRENT_MAKEFILE := Makefile"
echo "include unmake/Makefile"
echo "#-------------------------------------"
