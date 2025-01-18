import json
from functions.create_pg_conn import create_pg_conn
from functions.execute_stmt import execute_stmt

#region SQL Statements

insert_pg_stats = """
    INSERT INTO cfb.stats (
    team_id, season, interception_yards, turnovers, punt_return_yards, fumbles_lost, possession_time, 
    fumbles_recovered, rushing_tds, kick_return_yards, pass_attempts, interception_tds, 
    fourth_down_conversions, punt_returns, interceptions, third_down_conversions, passing_tds, 
    rushing_yards, penalty_yards, kick_return_tds, fourth_downs, penalties, third_downs, kick_returns, 
    total_yards, passes_intercepted, sacks, punt_return_tds, tackles_for_loss, pass_completions, 
    first_downs, rushing_attempts, net_passing_yards, o_plays, o_drives, o_ppa, o_total_ppa, 
    o_explosiveness, o_power_success, o_line_yards, o_line_yards_total, o_second_level_yards, 
    o_second_level_yards_total, o_open_field_yards, o_open_field_yards_total, o_avg_field_position_start, 
    o_avg_field_position_predicted_points, o_points_per_opportunity, o_total_opportunities, 
    d_plays, d_drives, d_ppa, d_total_ppa, d_explosiveness, d_power_success, d_line_yards, 
    d_line_yards_total, d_second_level_yards, d_second_level_yards_total, d_open_field_yards, 
    d_open_field_yards_total, d_avg_field_position_start, d_avg_field_position_predicted_points, 
    d_points_per_opportunity, d_total_opportunities
    ) 
    VALUES %s
    ON CONFLICT (team_id, season) DO NOTHING;
"""

select_pg_fbs_teams = """
    SELECT id, school
    FROM cfb.teams
    WHERE cfb.teams.classification = 'fbs';
"""

#endregion

try:
    # Set up db connection
    conn, cur = create_pg_conn()

    # Fetch all FBS schools with ids
    execute_stmt(cur, conn, select_pg_fbs_teams)
    rows = cur.fetchall()

    # Create mapping of school names to ids
    teams = {team[1]: {"id": team[0]} for team in rows}

    #region Get basic stats
    with open("data/basic_stats.json", "r") as file:
        data = json.load(file)["basic_stats"]
    
    for stat in data:
        teams[stat["team"]][stat["statName"]] = stat["statValue"]

    #endregion

    #region Get advanced stats
    with open("data/advanced_stats.json", "r") as file:
        data = json.load(file)["advanced_stats"]
    
    stats = []

    for stat in data:
        team = stat["team"]

        stats.append((
            teams[team]["id"],
            2024,
            teams[team]["interceptionYards"],
            teams[team]["turnovers"],
            teams[team]["puntReturnYards"],
            teams[team]["fumblesLost"],
            teams[team]["possessionTime"],
            teams[team]["fumblesRecovered"],
            teams[team]["rushingTDs"],
            teams[team]["kickReturnYards"],
            teams[team]["passAttempts"],
            teams[team]["interceptionTDs"],
            teams[team]["fourthDownConversions"],
            teams[team]["puntReturns"],
            teams[team]["interceptions"],
            teams[team]["thirdDownConversions"],
            teams[team]["passingTDs"],
            teams[team]["rushingYards"],
            teams[team]["penaltyYards"],
            teams[team]["kickReturnTDs"],
            teams[team]["fourthDowns"],
            teams[team]["penalties"],
            teams[team]["thirdDowns"],
            teams[team]["kickReturns"],
            teams[team]["totalYards"],
            teams[team]["passesIntercepted"],
            teams[team]["sacks"],
            teams[team]["puntReturnTDs"],
            teams[team]["tacklesForLoss"],
            teams[team]["passCompletions"],
            teams[team]["firstDowns"],
            teams[team]["rushingAttempts"],
            teams[team]["netPassingYards"],
            stat["offense"]["plays"],
            stat["offense"]["drives"],
            stat["offense"]["ppa"],
            stat["offense"]["totalPPA"],
            stat["offense"]["explosiveness"],
            stat["offense"]["powerSuccess"],
            stat["offense"]["lineYards"],
            stat["offense"]["lineYardsTotal"],
            stat["offense"]["secondLevelYards"],
            stat["offense"]["secondLevelYardsTotal"],
            stat["offense"]["openFieldYards"],
            stat["offense"]["openFieldYardsTotal"],
            stat["offense"]["fieldPosition"]["averageStart"],
            stat["offense"]["fieldPosition"]["averagePredictedPoints"],
            stat["offense"]["pointsPerOpportunity"],
            stat["offense"]["totalOpportunies"],
            stat["defense"]["plays"],
            stat["defense"]["drives"],
            stat["defense"]["ppa"],
            stat["defense"]["totalPPA"],
            stat["defense"]["explosiveness"],
            stat["defense"]["powerSuccess"],
            stat["defense"]["lineYards"],
            stat["defense"]["lineYardsTotal"],
            stat["defense"]["secondLevelYards"],
            stat["defense"]["secondLevelYardsTotal"],
            stat["defense"]["openFieldYards"],
            stat["defense"]["openFieldYardsTotal"],
            stat["defense"]["fieldPosition"]["averageStart"],
            stat["defense"]["fieldPosition"]["averagePredictedPoints"],
            stat["defense"]["pointsPerOpportunity"],
            stat["defense"]["totalOpportunies"]
        ))

    # Insert stats
    execute_stmt(cur, conn, insert_pg_stats, stats)

    print("Inserted stats")

    #endregion

    # Close db connection
    conn.close()
    cur.close()

except Exception as e:
    print(e.stack)