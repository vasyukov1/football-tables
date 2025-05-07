// frontend/src/api/group.ts
import apiClient from './client';
import { Group } from '../types/models';

export interface CreateGroupDto {
  name: string;
  teamIds: number[];
}

export const fetchGroups = () => apiClient.get<Group[]>('/groups');
export const createGroup = (data: CreateGroupDto) => apiClient.post('/groups', data);
