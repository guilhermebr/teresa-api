// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import "github.com/spf13/cobra"

// create a team
var teamCmd = &cobra.Command{
	Use:   "team",
	Short: "Create a team",
	Long:  `Create a team that is able to have many applications.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if teamNameFlag == "" {
			return newInputError("team name is required")
		}
		tc := NewTeresa()
		team, err := tc.CreateTeam(teamNameFlag, teamEmailFlag, teamURLFlag)
		if err != nil {
			log.Fatalf("Failed to create team: %s", err)
		}
		log.Infof("Team created. Name: %s Email: %s URL: %s\n", *team.Name, team.Email, team.URL)
		return err
	},
}

// delete team
var deleteTeamCmd = &cobra.Command{
	Use:   "team",
	Short: "Delete a team",
	Long:  `Delete a team`,
	Run: func(cmd *cobra.Command, args []string) {
		if teamIDFlag == 0 {
			Fatalf(cmd, "team ID is required")
		}
		if err := NewTeresa().DeleteTeam(teamIDFlag); err != nil {
			log.Fatalf("Failed to delete team: %s", err)
		}
		log.Infof("Team deleted.")
	},
}

func init() {
	createCmd.AddCommand(teamCmd)
	teamCmd.Flags().StringVarP(&teamNameFlag, "name", "n", "", "team name [required]")
	teamCmd.Flags().StringVarP(&teamEmailFlag, "email", "e", "", "team email, if any")
	teamCmd.Flags().StringVarP(&teamURLFlag, "url", "u", "", "team site's URL, if any")

	deleteCmd.AddCommand(deleteTeamCmd)
	deleteTeamCmd.Flags().Int64Var(&teamIDFlag, "id", 0, "team ID [required]")
}