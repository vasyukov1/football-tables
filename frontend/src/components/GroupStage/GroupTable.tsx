import React from 'react';
import { Team } from '../../types/models';
import { Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Paper, Typography } from '@mui/material';
import TeamStatRow from '../Shared/TeamStatRow';

const GroupTable: React.FC<{ teams: Team[]; groupName: string }> = ({ teams, groupName }) => {
    return (
        <TableContainer component={Paper} sx={{ mb: 4 }}>
            <Typography variant="h6" p={2}>{groupName}</Typography>
            <Table>
                <TableHead>
                    <TableRow>
                        <TableCell>Team</TableCell>
                        <TableCell align="right">Played</TableCell>
                        <TableCell align="right">W</TableCell>
                        <TableCell align="right">D</TableCell>
                        <TableCell align="right">L</TableCell>
                        <TableCell align="right">GF</TableCell>
                        <TableCell align="right">GA</TableCell>
                        <TableCell align="right">GD</TableCell>
                        <TableCell align="right">Pts</TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    {teams.map(team => (
                        <TeamStatRow key={team.id} team={team} />
                    ))}
                </TableBody>
            </Table>
        </TableContainer>
    );
};

export default GroupTable;