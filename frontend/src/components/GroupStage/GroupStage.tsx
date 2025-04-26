// src/components/GroupStage/GroupStage.tsx
import React from 'react';
import { Group } from '../../types/models';
import GroupTable from './GroupTable';
import { Box, Typography } from '@mui/material';

interface GroupStageProps {
    groups: Group[];
}

const GroupStage: React.FC<GroupStageProps> = ({ groups }) => {
    return (
        <Box sx={{ mt: 4 }}>
            <Typography variant="h4" gutterBottom sx={{ mb: 3 }}>
                Group Stage
            </Typography>

            <Box
                display="grid"
                gridTemplateColumns="repeat(auto-fit, minmax(300px, 1fr))"
                gap={4}
            >
                {groups.map(group => (
                    <GroupTable
                        key={group.name}
                        teams={group.teams}
                        groupName={group.name}
                    />
                ))}
            </Box>
        </Box>
    );
};

export default GroupStage;