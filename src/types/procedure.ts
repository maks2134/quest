export interface Procedure {
  id: number;
  title: string;
  type: string;
  content: string;
  sort_order: number;
  is_expanded: boolean;
  created_at: string;
  updated_at: string;
}

export interface ProcedureContentItem {
  type?: string;
  value: string;
  url?: string;
}

