import { gql } from '@apollo/client';
import * as Apollo from '@apollo/client';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
const defaultOptions = {} as const;
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
};

export type Mutation = {
  __typename?: 'Mutation';
  _empty?: Maybe<Scalars['String']>;
  createTodo: Todo;
  createUser: User;
};


export type MutationCreateTodoArgs = {
  input: NewTodo;
};


export type MutationCreateUserArgs = {
  input: NewUser;
};

export type NewTodo = {
  text: Scalars['String'];
  userId: Scalars['String'];
};

export type NewUser = {
  name: Scalars['String'];
};

export type Query = {
  __typename?: 'Query';
  _empty?: Maybe<Scalars['String']>;
  todos: Array<Todo>;
  users: Array<User>;
};

export type Todo = {
  __typename?: 'Todo';
  done: Scalars['Boolean'];
  id: Scalars['ID'];
  text: Scalars['String'];
  user: User;
};

export type User = {
  __typename?: 'User';
  id: Scalars['ID'];
  name: Scalars['String'];
};

export type GetTodosQueryVariables = Exact<{ [key: string]: never; }>;


export type GetTodosQuery = { __typename?: 'Query', todos: Array<{ __typename?: 'Todo', id: string, text: string, done: boolean }> };

export type AddTodoMutationVariables = Exact<{
  text: Scalars['String'];
  userId: Scalars['String'];
}>;


export type AddTodoMutation = { __typename?: 'Mutation', createTodo: { __typename?: 'Todo', text: string } };

export type AddUserMutationVariables = Exact<{
  name: Scalars['String'];
}>;


export type AddUserMutation = { __typename?: 'Mutation', createUser: { __typename?: 'User', name: string } };


export const GetTodosDocument = gql`
    query GetTodos {
  todos {
    id
    text
    done
  }
}
    `;

/**
 * __useGetTodosQuery__
 *
 * To run a query within a React component, call `useGetTodosQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetTodosQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetTodosQuery({
 *   variables: {
 *   },
 * });
 */
export function useGetTodosQuery(baseOptions?: Apollo.QueryHookOptions<GetTodosQuery, GetTodosQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<GetTodosQuery, GetTodosQueryVariables>(GetTodosDocument, options);
      }
export function useGetTodosLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<GetTodosQuery, GetTodosQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<GetTodosQuery, GetTodosQueryVariables>(GetTodosDocument, options);
        }
export type GetTodosQueryHookResult = ReturnType<typeof useGetTodosQuery>;
export type GetTodosLazyQueryHookResult = ReturnType<typeof useGetTodosLazyQuery>;
export type GetTodosQueryResult = Apollo.QueryResult<GetTodosQuery, GetTodosQueryVariables>;
export const AddTodoDocument = gql`
    mutation AddTodo($text: String!, $userId: String!) {
  createTodo(input: {text: $text, userId: $userId}) {
    text
  }
}
    `;
export type AddTodoMutationFn = Apollo.MutationFunction<AddTodoMutation, AddTodoMutationVariables>;

/**
 * __useAddTodoMutation__
 *
 * To run a mutation, you first call `useAddTodoMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useAddTodoMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [addTodoMutation, { data, loading, error }] = useAddTodoMutation({
 *   variables: {
 *      text: // value for 'text'
 *      userId: // value for 'userId'
 *   },
 * });
 */
export function useAddTodoMutation(baseOptions?: Apollo.MutationHookOptions<AddTodoMutation, AddTodoMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<AddTodoMutation, AddTodoMutationVariables>(AddTodoDocument, options);
      }
export type AddTodoMutationHookResult = ReturnType<typeof useAddTodoMutation>;
export type AddTodoMutationResult = Apollo.MutationResult<AddTodoMutation>;
export type AddTodoMutationOptions = Apollo.BaseMutationOptions<AddTodoMutation, AddTodoMutationVariables>;
export const AddUserDocument = gql`
    mutation AddUser($name: String!) {
  createUser(input: {name: $name}) {
    name
  }
}
    `;
export type AddUserMutationFn = Apollo.MutationFunction<AddUserMutation, AddUserMutationVariables>;

/**
 * __useAddUserMutation__
 *
 * To run a mutation, you first call `useAddUserMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useAddUserMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [addUserMutation, { data, loading, error }] = useAddUserMutation({
 *   variables: {
 *      name: // value for 'name'
 *   },
 * });
 */
export function useAddUserMutation(baseOptions?: Apollo.MutationHookOptions<AddUserMutation, AddUserMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<AddUserMutation, AddUserMutationVariables>(AddUserDocument, options);
      }
export type AddUserMutationHookResult = ReturnType<typeof useAddUserMutation>;
export type AddUserMutationResult = Apollo.MutationResult<AddUserMutation>;
export type AddUserMutationOptions = Apollo.BaseMutationOptions<AddUserMutation, AddUserMutationVariables>;