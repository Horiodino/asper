package cmd

func requiredflags(){
	// input should be like this go run main.go createvms --name vmname --os fedora-server --cpu 4 --memory 3Gi --disk 20

	createvms.Flags().StringVarP(&name, "name", "n", "", "Name of the Virtual Machine")
	createvms.MarkFlagRequired("name")
	createvms.Flags().StringVarP(&vmos, "os", "o", "", "Operating System of the Virtual Machine")
	createvms.MarkFlagRequired("os")
	createvms.Flags().StringVarP(&cpu, "cpu", "c", "", "Number of CPU cores")
	createvms.MarkFlagRequired("cpu")
	createvms.Flags().StringVarP(&memory, "memory", "m", "", "Memory of the Virtual Machine")
	createvms.MarkFlagRequired("memory")
	createvms.Flags().StringVarP(&disk, "disk", "d", "", "Disk size of the Virtual Machine")
	createvms.MarkFlagRequired("disk")
	
}