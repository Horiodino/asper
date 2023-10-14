package cmd

func requiredflags(){
	// input should be like this go run main.go createvm --name vmname --os fedora-server --cpu 4 --memory 3Gi --disk 20

	createvm.Flags().StringVarP(&name, "name", "n", "", "Name of the Virtual Machine")
	createvm.MarkFlagRequired("name")
	createvm.Flags().StringVarP(&vmos, "os", "o", "", "Operating System of the Virtual Machine")
	createvm.MarkFlagRequired("os")
	createvm.Flags().StringVarP(&cpu, "cpu", "c", "", "Number of CPU cores")
	createvm.MarkFlagRequired("cpu")
	createvm.Flags().StringVarP(&memory, "memory", "m", "", "Memory of the Virtual Machine")
	createvm.MarkFlagRequired("memory")
	createvm.Flags().StringVarP(&disk, "disk", "d", "", "Disk size of the Virtual Machine")
	createvm.MarkFlagRequired("disk")
	
}