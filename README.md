Let do an extensible cli with my explaination and others here.


rootCmd.AddCommand(net.NetCmd): we use to add a new command to root when we make the var public

It is just like we create a command first and have subsidiary under it, for instance now, we have net the dada with ping as the subsidiary, and net is always public while ping is private and can be called from under it and also we have info and we can do disk-checking in it


get viper from cobra with (cobra-cli init --viper)
We can use viper to call datatypes with their values ot the cli and necessary use it also


TODO:

Make it in such a way the cmd and the help with et all are more acessible and easy to use
