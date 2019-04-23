package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var rootCmd = &cobra.Command{
	Use:   "redeploy {namespace} {service}",
	Args:  cobra.ExactArgs(2),
	Short: "Redeploys a service in Rancher",
	Run: func(cmd *cobra.Command, args []string) {
		// create the clientset
		clientset, err := kubernetes.NewForConfig(getConfig(cmd))
		if err != nil {
			panic(err.Error())
		}

		dep, err := clientset.AppsV1beta2().Deployments(args[0]).Get(args[1], metav1.GetOptions{})
		if err != nil {
			panic(err.Error())
		}

		rev, _ := strconv.Atoi(dep.Annotations["deployment.kubernetes.io/revision"])
		dep.Annotations["deployment.kubernetes.io/revision"] = strconv.Itoa(rev + 1)
		dep.Spec.Template.Annotations["cattle.io/timestamp"] = time.Now().Format(time.RFC3339)

		if _, err := clientset.AppsV1beta2().Deployments(args[0]).Update(dep); err != nil {
			panic(err.Error())
		}

		fmt.Fprintf(os.Stdout, "Service %s:%s redeployed!\n", args[0], args[1])
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var development bool
var kubeconfig string

func init() {
	rootCmd.Flags().BoolVarP(&development, "development", "d", false, "Use out-of-cluster config")
	rootCmd.Flags().StringVarP(&kubeconfig, "kubeconfig", "c", filepath.Join(os.Getenv("HOME"), ".kube", "config"), "kubeconfig file to use while using out-of-cluster config")
}
