package cmd

import (
	"github.com/dendi239/gitfame/pkg/filter"
	gitfame "github.com/dendi239/gitfame/pkg/git"
	"github.com/dendi239/gitfame/pkg/printer"
	"github.com/spf13/cobra"
)

var (
	g = gitfame.NewContext(f.CanTake)
	f filter.Filter
	p printer.Printer
)

var rootCmd = &cobra.Command{
	Use:   "gitfame",
	Short: "Gitfame counts some statistics over git-repo",
	Long:  `gitfame calculates (lines, files, commits) for selected revision`,
	Run: func(cmd *cobra.Command, args []string) {
		authors, err := g.Gitfame()
		if err != nil {
			panic(err)
		}
		p.Print(authors)
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringVar(&g.Directory, "repository", ".", "path to repository")
	rootCmd.PersistentFlags().StringVar(&g.Revision, "revision", "HEAD", "revision to analyze")
	rootCmd.PersistentFlags().StringVar(&p.OrderBy, "order-by", "lines", `key for result sorting, one of: "lines", "commits", "files"`)
	rootCmd.PersistentFlags().StringVar(&p.Format, "format", "tabular", `format, one of: "tabular", "json", "csv", "json", "json-lines"`)
	rootCmd.PersistentFlags().StringSliceVar(&f.Extensions, "extensions", []string{}, "extensions to count, e.g. '.go,.md'")
	rootCmd.PersistentFlags().StringSliceVar(&f.Languages, "languages", []string{}, "languages to count, e.g. 'go,markdown'")
	rootCmd.PersistentFlags().StringSliceVar(&f.Exclude, "exclude", []string{}, "glob patterns to exclude from counting")
	rootCmd.PersistentFlags().StringSliceVar(&f.RestrictTo, "restrict-to", []string{}, "if specified, every counted file should match at least one of these globs")
	rootCmd.PersistentFlags().BoolVar(&g.UseCommiter, "use-committer", false, "Using committer instead of author")
	rootCmd.PersistentFlags().BoolVarP(&g.Progress, "progress", "p", false, "shows actual progress during analyzing")
}
