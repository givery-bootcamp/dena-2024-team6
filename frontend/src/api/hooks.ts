/**
 * Generated by orval v6.30.2 🍺
 * Do not edit manually.
 * Web開発研修6班 API
 * FY24卒Web開発研修6班のAPI仕様書です
 * OpenAPI spec version: 1.0.0
 */
import {
  useMutation,
  useQuery
} from '@tanstack/react-query'
import type {
  MutationFunction,
  QueryFunction,
  QueryKey,
  UseMutationOptions,
  UseMutationResult,
  UseQueryOptions,
  UseQueryResult
} from '@tanstack/react-query'
import type {
  SchemaCommentResponse,
  SchemaCreateCommentRequest,
  SchemaCreatePostRequest,
  SchemaErrorResponse,
  SchemaLikeRecordResponse,
  SchemaLoginRequest,
  SchemaMutationSchema,
  SchemaPostDetailResponse,
  SchemaPostResponse,
  SchemaSignupRequest,
  SchemaSpeedResponse,
  SchemaUpdateCommentRequest,
  SchemaUpdatePostRequest,
  SchemaUserResponse
} from './model'
import { customInstance } from '../shared/libs/axios';


type SecondParameter<T extends (...args: any) => any> = Parameters<T>[1];


/**
 * @summary APIのセルフチェック
 */
export const healthCheck = (
    
 options?: SecondParameter<typeof customInstance>,signal?: AbortSignal
) => {
      
      
      return customInstance<void>(
      {url: `/`, method: 'GET', signal
    },
      options);
    }
  

export const getHealthCheckQueryKey = () => {
    return [`/`] as const;
    }

    
export const getHealthCheckQueryOptions = <TData = Awaited<ReturnType<typeof healthCheck>>, TError = unknown>( options?: { query?:Partial<UseQueryOptions<Awaited<ReturnType<typeof healthCheck>>, TError, TData>>, request?: SecondParameter<typeof customInstance>}
) => {

const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey =  queryOptions?.queryKey ?? getHealthCheckQueryKey();

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof healthCheck>>> = ({ signal }) => healthCheck(requestOptions, signal);

      

      

   return  { queryKey, queryFn, ...queryOptions} as UseQueryOptions<Awaited<ReturnType<typeof healthCheck>>, TError, TData> & { queryKey: QueryKey }
}

export type HealthCheckQueryResult = NonNullable<Awaited<ReturnType<typeof healthCheck>>>
export type HealthCheckQueryError = unknown

/**
 * @summary APIのセルフチェック
 */
export const useHealthCheck = <TData = Awaited<ReturnType<typeof healthCheck>>, TError = unknown>(
  options?: { query?:Partial<UseQueryOptions<Awaited<ReturnType<typeof healthCheck>>, TError, TData>>, request?: SecondParameter<typeof customInstance>}

  ):  UseQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getHealthCheckQueryOptions(options)

  const query = useQuery(queryOptions) as  UseQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryOptions.queryKey ;

  return query;
}




/**
 * @summary 投稿の一覧を取得
 */
export const listPosts = (
    
 options?: SecondParameter<typeof customInstance>,signal?: AbortSignal
) => {
      
      
      return customInstance<SchemaPostResponse[]>(
      {url: `/posts`, method: 'GET', signal
    },
      options);
    }
  

export const getListPostsQueryKey = () => {
    return [`/posts`] as const;
    }

    
export const getListPostsQueryOptions = <TData = Awaited<ReturnType<typeof listPosts>>, TError = SchemaErrorResponse>( options?: { query?:Partial<UseQueryOptions<Awaited<ReturnType<typeof listPosts>>, TError, TData>>, request?: SecondParameter<typeof customInstance>}
) => {

const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey =  queryOptions?.queryKey ?? getListPostsQueryKey();

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof listPosts>>> = ({ signal }) => listPosts(requestOptions, signal);

      

      

   return  { queryKey, queryFn, ...queryOptions} as UseQueryOptions<Awaited<ReturnType<typeof listPosts>>, TError, TData> & { queryKey: QueryKey }
}

export type ListPostsQueryResult = NonNullable<Awaited<ReturnType<typeof listPosts>>>
export type ListPostsQueryError = SchemaErrorResponse

/**
 * @summary 投稿の一覧を取得
 */
export const useListPosts = <TData = Awaited<ReturnType<typeof listPosts>>, TError = SchemaErrorResponse>(
  options?: { query?:Partial<UseQueryOptions<Awaited<ReturnType<typeof listPosts>>, TError, TData>>, request?: SecondParameter<typeof customInstance>}

  ):  UseQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getListPostsQueryOptions(options)

  const query = useQuery(queryOptions) as  UseQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryOptions.queryKey ;

  return query;
}




/**
 * @summary 投稿を作成
 */
export const createPost = (
    schemaCreatePostRequest: SchemaCreatePostRequest,
 options?: SecondParameter<typeof customInstance>,) => {
      
      
      return customInstance<SchemaPostResponse[]>(
      {url: `/posts`, method: 'POST',
      headers: {'Content-Type': 'application/json', },
      data: schemaCreatePostRequest
    },
      options);
    }
  


export const getCreatePostMutationOptions = <TError = SchemaErrorResponse,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof createPost>>, TError,{data: SchemaCreatePostRequest}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationOptions<Awaited<ReturnType<typeof createPost>>, TError,{data: SchemaCreatePostRequest}, TContext> => {
const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof createPost>>, {data: SchemaCreatePostRequest}> = (props) => {
          const {data} = props ?? {};

          return  createPost(data,requestOptions)
        }

        


  return  { mutationFn, ...mutationOptions }}

    export type CreatePostMutationResult = NonNullable<Awaited<ReturnType<typeof createPost>>>
    export type CreatePostMutationBody = SchemaCreatePostRequest
    export type CreatePostMutationError = SchemaErrorResponse

    /**
 * @summary 投稿を作成
 */
export const useCreatePost = <TError = SchemaErrorResponse,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof createPost>>, TError,{data: SchemaCreatePostRequest}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationResult<
        Awaited<ReturnType<typeof createPost>>,
        TError,
        {data: SchemaCreatePostRequest},
        TContext
      > => {

      const mutationOptions = getCreatePostMutationOptions(options);

      return useMutation(mutationOptions);
    }
    
/**
 * @summary ユーザのログインを実行
 */
export const signIn = (
    schemaLoginRequest: SchemaLoginRequest,
 options?: SecondParameter<typeof customInstance>,) => {
      
      
      return customInstance<SchemaUserResponse>(
      {url: `/signin`, method: 'POST',
      headers: {'Content-Type': 'application/json', },
      data: schemaLoginRequest
    },
      options);
    }
  


export const getSignInMutationOptions = <TError = SchemaErrorResponse,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof signIn>>, TError,{data: SchemaLoginRequest}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationOptions<Awaited<ReturnType<typeof signIn>>, TError,{data: SchemaLoginRequest}, TContext> => {
const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof signIn>>, {data: SchemaLoginRequest}> = (props) => {
          const {data} = props ?? {};

          return  signIn(data,requestOptions)
        }

        


  return  { mutationFn, ...mutationOptions }}

    export type SignInMutationResult = NonNullable<Awaited<ReturnType<typeof signIn>>>
    export type SignInMutationBody = SchemaLoginRequest
    export type SignInMutationError = SchemaErrorResponse

    /**
 * @summary ユーザのログインを実行
 */
export const useSignIn = <TError = SchemaErrorResponse,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof signIn>>, TError,{data: SchemaLoginRequest}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationResult<
        Awaited<ReturnType<typeof signIn>>,
        TError,
        {data: SchemaLoginRequest},
        TContext
      > => {

      const mutationOptions = getSignInMutationOptions(options);

      return useMutation(mutationOptions);
    }
    
/**
 * @summary ユーザのログアウトを実行
 */
export const signOut = (
    
 options?: SecondParameter<typeof customInstance>,) => {
      
      
      return customInstance<unknown>(
      {url: `/signout`, method: 'POST'
    },
      options);
    }
  


export const getSignOutMutationOptions = <TError = SchemaErrorResponse,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof signOut>>, TError,void, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationOptions<Awaited<ReturnType<typeof signOut>>, TError,void, TContext> => {
const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof signOut>>, void> = () => {
          

          return  signOut(requestOptions)
        }

        


  return  { mutationFn, ...mutationOptions }}

    export type SignOutMutationResult = NonNullable<Awaited<ReturnType<typeof signOut>>>
    
    export type SignOutMutationError = SchemaErrorResponse

    /**
 * @summary ユーザのログアウトを実行
 */
export const useSignOut = <TError = SchemaErrorResponse,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof signOut>>, TError,void, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationResult<
        Awaited<ReturnType<typeof signOut>>,
        TError,
        void,
        TContext
      > => {

      const mutationOptions = getSignOutMutationOptions(options);

      return useMutation(mutationOptions);
    }
    
/**
 * @summary ユーザのアカウント登録を実行
 */
export const signUp = (
    schemaSignupRequest: SchemaSignupRequest,
 options?: SecondParameter<typeof customInstance>,) => {
      
      
      return customInstance<SchemaUserResponse>(
      {url: `/signup`, method: 'POST',
      headers: {'Content-Type': 'application/json', },
      data: schemaSignupRequest
    },
      options);
    }
  


export const getSignUpMutationOptions = <TError = SchemaErrorResponse,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof signUp>>, TError,{data: SchemaSignupRequest}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationOptions<Awaited<ReturnType<typeof signUp>>, TError,{data: SchemaSignupRequest}, TContext> => {
const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof signUp>>, {data: SchemaSignupRequest}> = (props) => {
          const {data} = props ?? {};

          return  signUp(data,requestOptions)
        }

        


  return  { mutationFn, ...mutationOptions }}

    export type SignUpMutationResult = NonNullable<Awaited<ReturnType<typeof signUp>>>
    export type SignUpMutationBody = SchemaSignupRequest
    export type SignUpMutationError = SchemaErrorResponse

    /**
 * @summary ユーザのアカウント登録を実行
 */
export const useSignUp = <TError = SchemaErrorResponse,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof signUp>>, TError,{data: SchemaSignupRequest}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationResult<
        Awaited<ReturnType<typeof signUp>>,
        TError,
        {data: SchemaSignupRequest},
        TContext
      > => {

      const mutationOptions = getSignUpMutationOptions(options);

      return useMutation(mutationOptions);
    }
    
/**
 * @summary 現在ログインしているユーザを取得
 */
export const getCurrentUser = (
    
 options?: SecondParameter<typeof customInstance>,signal?: AbortSignal
) => {
      
      
      return customInstance<SchemaUserResponse>(
      {url: `/user`, method: 'GET', signal
    },
      options);
    }
  

export const getGetCurrentUserQueryKey = () => {
    return [`/user`] as const;
    }

    
export const getGetCurrentUserQueryOptions = <TData = Awaited<ReturnType<typeof getCurrentUser>>, TError = SchemaErrorResponse>( options?: { query?:Partial<UseQueryOptions<Awaited<ReturnType<typeof getCurrentUser>>, TError, TData>>, request?: SecondParameter<typeof customInstance>}
) => {

const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey =  queryOptions?.queryKey ?? getGetCurrentUserQueryKey();

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof getCurrentUser>>> = ({ signal }) => getCurrentUser(requestOptions, signal);

      

      

   return  { queryKey, queryFn, ...queryOptions} as UseQueryOptions<Awaited<ReturnType<typeof getCurrentUser>>, TError, TData> & { queryKey: QueryKey }
}

export type GetCurrentUserQueryResult = NonNullable<Awaited<ReturnType<typeof getCurrentUser>>>
export type GetCurrentUserQueryError = SchemaErrorResponse

/**
 * @summary 現在ログインしているユーザを取得
 */
export const useGetCurrentUser = <TData = Awaited<ReturnType<typeof getCurrentUser>>, TError = SchemaErrorResponse>(
  options?: { query?:Partial<UseQueryOptions<Awaited<ReturnType<typeof getCurrentUser>>, TError, TData>>, request?: SecondParameter<typeof customInstance>}

  ):  UseQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getGetCurrentUserQueryOptions(options)

  const query = useQuery(queryOptions) as  UseQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryOptions.queryKey ;

  return query;
}




/**
 * @summary 投稿をIDから取得
 */
export const getPost = (
    id: string,
 options?: SecondParameter<typeof customInstance>,signal?: AbortSignal
) => {
      
      
      return customInstance<SchemaPostDetailResponse>(
      {url: `/posts/${id}`, method: 'GET', signal
    },
      options);
    }
  

export const getGetPostQueryKey = (id: string,) => {
    return [`/posts/${id}`] as const;
    }

    
export const getGetPostQueryOptions = <TData = Awaited<ReturnType<typeof getPost>>, TError = SchemaErrorResponse>(id: string, options?: { query?:Partial<UseQueryOptions<Awaited<ReturnType<typeof getPost>>, TError, TData>>, request?: SecondParameter<typeof customInstance>}
) => {

const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey =  queryOptions?.queryKey ?? getGetPostQueryKey(id);

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof getPost>>> = ({ signal }) => getPost(id, requestOptions, signal);

      

      

   return  { queryKey, queryFn, enabled: !!(id), ...queryOptions} as UseQueryOptions<Awaited<ReturnType<typeof getPost>>, TError, TData> & { queryKey: QueryKey }
}

export type GetPostQueryResult = NonNullable<Awaited<ReturnType<typeof getPost>>>
export type GetPostQueryError = SchemaErrorResponse

/**
 * @summary 投稿をIDから取得
 */
export const useGetPost = <TData = Awaited<ReturnType<typeof getPost>>, TError = SchemaErrorResponse>(
 id: string, options?: { query?:Partial<UseQueryOptions<Awaited<ReturnType<typeof getPost>>, TError, TData>>, request?: SecondParameter<typeof customInstance>}

  ):  UseQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getGetPostQueryOptions(id,options)

  const query = useQuery(queryOptions) as  UseQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryOptions.queryKey ;

  return query;
}




/**
 * @summary 対象の投稿のコメント一覧を取得
 */
export const listPostComments = (
    postId: string,
 options?: SecondParameter<typeof customInstance>,signal?: AbortSignal
) => {
      
      
      return customInstance<SchemaCommentResponse[]>(
      {url: `/posts/${postId}/comments`, method: 'GET', signal
    },
      options);
    }
  

export const getListPostCommentsQueryKey = (postId: string,) => {
    return [`/posts/${postId}/comments`] as const;
    }

    
export const getListPostCommentsQueryOptions = <TData = Awaited<ReturnType<typeof listPostComments>>, TError = SchemaErrorResponse>(postId: string, options?: { query?:Partial<UseQueryOptions<Awaited<ReturnType<typeof listPostComments>>, TError, TData>>, request?: SecondParameter<typeof customInstance>}
) => {

const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey =  queryOptions?.queryKey ?? getListPostCommentsQueryKey(postId);

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof listPostComments>>> = ({ signal }) => listPostComments(postId, requestOptions, signal);

      

      

   return  { queryKey, queryFn, enabled: !!(postId), ...queryOptions} as UseQueryOptions<Awaited<ReturnType<typeof listPostComments>>, TError, TData> & { queryKey: QueryKey }
}

export type ListPostCommentsQueryResult = NonNullable<Awaited<ReturnType<typeof listPostComments>>>
export type ListPostCommentsQueryError = SchemaErrorResponse

/**
 * @summary 対象の投稿のコメント一覧を取得
 */
export const useListPostComments = <TData = Awaited<ReturnType<typeof listPostComments>>, TError = SchemaErrorResponse>(
 postId: string, options?: { query?:Partial<UseQueryOptions<Awaited<ReturnType<typeof listPostComments>>, TError, TData>>, request?: SecondParameter<typeof customInstance>}

  ):  UseQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getListPostCommentsQueryOptions(postId,options)

  const query = useQuery(queryOptions) as  UseQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryOptions.queryKey ;

  return query;
}




/**
 * @summary 対象の投稿のコメントを追加
 */
export const createPostComments = (
    postId: string,
    schemaCreateCommentRequest: SchemaCreateCommentRequest,
 options?: SecondParameter<typeof customInstance>,) => {
      
      
      return customInstance<SchemaMutationSchema>(
      {url: `/posts/${postId}/comments`, method: 'POST',
      headers: {'Content-Type': 'application/json', },
      data: schemaCreateCommentRequest
    },
      options);
    }
  


export const getCreatePostCommentsMutationOptions = <TError = SchemaErrorResponse,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof createPostComments>>, TError,{postId: string;data: SchemaCreateCommentRequest}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationOptions<Awaited<ReturnType<typeof createPostComments>>, TError,{postId: string;data: SchemaCreateCommentRequest}, TContext> => {
const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof createPostComments>>, {postId: string;data: SchemaCreateCommentRequest}> = (props) => {
          const {postId,data} = props ?? {};

          return  createPostComments(postId,data,requestOptions)
        }

        


  return  { mutationFn, ...mutationOptions }}

    export type CreatePostCommentsMutationResult = NonNullable<Awaited<ReturnType<typeof createPostComments>>>
    export type CreatePostCommentsMutationBody = SchemaCreateCommentRequest
    export type CreatePostCommentsMutationError = SchemaErrorResponse

    /**
 * @summary 対象の投稿のコメントを追加
 */
export const useCreatePostComments = <TError = SchemaErrorResponse,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof createPostComments>>, TError,{postId: string;data: SchemaCreateCommentRequest}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationResult<
        Awaited<ReturnType<typeof createPostComments>>,
        TError,
        {postId: string;data: SchemaCreateCommentRequest},
        TContext
      > => {

      const mutationOptions = getCreatePostCommentsMutationOptions(options);

      return useMutation(mutationOptions);
    }
    
/**
 * @summary 対象の投稿のコメントを削除
 */
export const deletePostComments = (
    postId: string,
    commentId: string,
 options?: SecondParameter<typeof customInstance>,) => {
      
      
      return customInstance<SchemaMutationSchema>(
      {url: `/posts/${postId}/comments/${commentId}`, method: 'DELETE'
    },
      options);
    }
  


export const getDeletePostCommentsMutationOptions = <TError = SchemaErrorResponse,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof deletePostComments>>, TError,{postId: string;commentId: string}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationOptions<Awaited<ReturnType<typeof deletePostComments>>, TError,{postId: string;commentId: string}, TContext> => {
const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof deletePostComments>>, {postId: string;commentId: string}> = (props) => {
          const {postId,commentId} = props ?? {};

          return  deletePostComments(postId,commentId,requestOptions)
        }

        


  return  { mutationFn, ...mutationOptions }}

    export type DeletePostCommentsMutationResult = NonNullable<Awaited<ReturnType<typeof deletePostComments>>>
    
    export type DeletePostCommentsMutationError = SchemaErrorResponse

    /**
 * @summary 対象の投稿のコメントを削除
 */
export const useDeletePostComments = <TError = SchemaErrorResponse,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof deletePostComments>>, TError,{postId: string;commentId: string}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationResult<
        Awaited<ReturnType<typeof deletePostComments>>,
        TError,
        {postId: string;commentId: string},
        TContext
      > => {

      const mutationOptions = getDeletePostCommentsMutationOptions(options);

      return useMutation(mutationOptions);
    }
    
/**
 * @summary 対象の投稿のコメントを変更
 */
export const putPostComments = (
    postId: string,
    commentId: string,
    schemaUpdateCommentRequest: SchemaUpdateCommentRequest,
 options?: SecondParameter<typeof customInstance>,) => {
      
      
      return customInstance<SchemaMutationSchema>(
      {url: `/posts/${postId}/comments/${commentId}`, method: 'POST',
      headers: {'Content-Type': 'application/json', },
      data: schemaUpdateCommentRequest
    },
      options);
    }
  


export const getPutPostCommentsMutationOptions = <TError = SchemaErrorResponse,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof putPostComments>>, TError,{postId: string;commentId: string;data: SchemaUpdateCommentRequest}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationOptions<Awaited<ReturnType<typeof putPostComments>>, TError,{postId: string;commentId: string;data: SchemaUpdateCommentRequest}, TContext> => {
const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof putPostComments>>, {postId: string;commentId: string;data: SchemaUpdateCommentRequest}> = (props) => {
          const {postId,commentId,data} = props ?? {};

          return  putPostComments(postId,commentId,data,requestOptions)
        }

        


  return  { mutationFn, ...mutationOptions }}

    export type PutPostCommentsMutationResult = NonNullable<Awaited<ReturnType<typeof putPostComments>>>
    export type PutPostCommentsMutationBody = SchemaUpdateCommentRequest
    export type PutPostCommentsMutationError = SchemaErrorResponse

    /**
 * @summary 対象の投稿のコメントを変更
 */
export const usePutPostComments = <TError = SchemaErrorResponse,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof putPostComments>>, TError,{postId: string;commentId: string;data: SchemaUpdateCommentRequest}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationResult<
        Awaited<ReturnType<typeof putPostComments>>,
        TError,
        {postId: string;commentId: string;data: SchemaUpdateCommentRequest},
        TContext
      > => {

      const mutationOptions = getPutPostCommentsMutationOptions(options);

      return useMutation(mutationOptions);
    }
    
/**
 * @summary 対象の投稿のlike数を取得する
 */
export const getlikeRecord = (
    postId: string,
 options?: SecondParameter<typeof customInstance>,signal?: AbortSignal
) => {
      
      
      return customInstance<SchemaLikeRecordResponse>(
      {url: `/posts/${postId}/like`, method: 'GET', signal
    },
      options);
    }
  

export const getGetlikeRecordQueryKey = (postId: string,) => {
    return [`/posts/${postId}/like`] as const;
    }

    
export const getGetlikeRecordQueryOptions = <TData = Awaited<ReturnType<typeof getlikeRecord>>, TError = SchemaErrorResponse>(postId: string, options?: { query?:Partial<UseQueryOptions<Awaited<ReturnType<typeof getlikeRecord>>, TError, TData>>, request?: SecondParameter<typeof customInstance>}
) => {

const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey =  queryOptions?.queryKey ?? getGetlikeRecordQueryKey(postId);

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof getlikeRecord>>> = ({ signal }) => getlikeRecord(postId, requestOptions, signal);

      

      

   return  { queryKey, queryFn, enabled: !!(postId), ...queryOptions} as UseQueryOptions<Awaited<ReturnType<typeof getlikeRecord>>, TError, TData> & { queryKey: QueryKey }
}

export type GetlikeRecordQueryResult = NonNullable<Awaited<ReturnType<typeof getlikeRecord>>>
export type GetlikeRecordQueryError = SchemaErrorResponse

/**
 * @summary 対象の投稿のlike数を取得する
 */
export const useGetlikeRecord = <TData = Awaited<ReturnType<typeof getlikeRecord>>, TError = SchemaErrorResponse>(
 postId: string, options?: { query?:Partial<UseQueryOptions<Awaited<ReturnType<typeof getlikeRecord>>, TError, TData>>, request?: SecondParameter<typeof customInstance>}

  ):  UseQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getGetlikeRecordQueryOptions(postId,options)

  const query = useQuery(queryOptions) as  UseQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryOptions.queryKey ;

  return query;
}




/**
 * @summary 対象の投稿をlikeする
 */
export const likePost = (
    postId: string,
 options?: SecondParameter<typeof customInstance>,) => {
      
      
      return customInstance<SchemaMutationSchema>(
      {url: `/posts/${postId}/like`, method: 'POST'
    },
      options);
    }
  


export const getLikePostMutationOptions = <TError = SchemaErrorResponse,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof likePost>>, TError,{postId: string}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationOptions<Awaited<ReturnType<typeof likePost>>, TError,{postId: string}, TContext> => {
const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof likePost>>, {postId: string}> = (props) => {
          const {postId} = props ?? {};

          return  likePost(postId,requestOptions)
        }

        


  return  { mutationFn, ...mutationOptions }}

    export type LikePostMutationResult = NonNullable<Awaited<ReturnType<typeof likePost>>>
    
    export type LikePostMutationError = SchemaErrorResponse

    /**
 * @summary 対象の投稿をlikeする
 */
export const useLikePost = <TError = SchemaErrorResponse,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof likePost>>, TError,{postId: string}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationResult<
        Awaited<ReturnType<typeof likePost>>,
        TError,
        {postId: string},
        TContext
      > => {

      const mutationOptions = getLikePostMutationOptions(options);

      return useMutation(mutationOptions);
    }
    
/**
 * @summary 投稿を削除
 */
export const deletePost = (
    postid: string,
 options?: SecondParameter<typeof customInstance>,) => {
      
      
      return customInstance<SchemaMutationSchema>(
      {url: `/posts/${postid}`, method: 'DELETE'
    },
      options);
    }
  


export const getDeletePostMutationOptions = <TError = SchemaErrorResponse,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof deletePost>>, TError,{postid: string}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationOptions<Awaited<ReturnType<typeof deletePost>>, TError,{postid: string}, TContext> => {
const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof deletePost>>, {postid: string}> = (props) => {
          const {postid} = props ?? {};

          return  deletePost(postid,requestOptions)
        }

        


  return  { mutationFn, ...mutationOptions }}

    export type DeletePostMutationResult = NonNullable<Awaited<ReturnType<typeof deletePost>>>
    
    export type DeletePostMutationError = SchemaErrorResponse

    /**
 * @summary 投稿を削除
 */
export const useDeletePost = <TError = SchemaErrorResponse,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof deletePost>>, TError,{postid: string}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationResult<
        Awaited<ReturnType<typeof deletePost>>,
        TError,
        {postid: string},
        TContext
      > => {

      const mutationOptions = getDeletePostMutationOptions(options);

      return useMutation(mutationOptions);
    }
    
/**
 * @summary 投稿を更新
 */
export const updatePost = (
    postid: string,
    schemaUpdatePostRequest: SchemaUpdatePostRequest,
 options?: SecondParameter<typeof customInstance>,) => {
      
      
      return customInstance<SchemaMutationSchema>(
      {url: `/posts/${postid}`, method: 'PUT',
      headers: {'Content-Type': 'application/json', },
      data: schemaUpdatePostRequest
    },
      options);
    }
  


export const getUpdatePostMutationOptions = <TError = SchemaErrorResponse,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof updatePost>>, TError,{postid: string;data: SchemaUpdatePostRequest}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationOptions<Awaited<ReturnType<typeof updatePost>>, TError,{postid: string;data: SchemaUpdatePostRequest}, TContext> => {
const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof updatePost>>, {postid: string;data: SchemaUpdatePostRequest}> = (props) => {
          const {postid,data} = props ?? {};

          return  updatePost(postid,data,requestOptions)
        }

        


  return  { mutationFn, ...mutationOptions }}

    export type UpdatePostMutationResult = NonNullable<Awaited<ReturnType<typeof updatePost>>>
    export type UpdatePostMutationBody = SchemaUpdatePostRequest
    export type UpdatePostMutationError = SchemaErrorResponse

    /**
 * @summary 投稿を更新
 */
export const useUpdatePost = <TError = SchemaErrorResponse,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof updatePost>>, TError,{postid: string;data: SchemaUpdatePostRequest}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationResult<
        Awaited<ReturnType<typeof updatePost>>,
        TError,
        {postid: string;data: SchemaUpdatePostRequest},
        TContext
      > => {

      const mutationOptions = getUpdatePostMutationOptions(options);

      return useMutation(mutationOptions);
    }
    
/**
 * @summary 定期的に実行される。1h経過したらlike数を更新する
 */
export const updatelikeRecord = (
    
 options?: SecondParameter<typeof customInstance>,) => {
      
      
      return customInstance<SchemaLikeRecordResponse>(
      {url: `/posts/like/update`, method: 'POST'
    },
      options);
    }
  


export const getUpdatelikeRecordMutationOptions = <TError = SchemaErrorResponse,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof updatelikeRecord>>, TError,void, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationOptions<Awaited<ReturnType<typeof updatelikeRecord>>, TError,void, TContext> => {
const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof updatelikeRecord>>, void> = () => {
          

          return  updatelikeRecord(requestOptions)
        }

        


  return  { mutationFn, ...mutationOptions }}

    export type UpdatelikeRecordMutationResult = NonNullable<Awaited<ReturnType<typeof updatelikeRecord>>>
    
    export type UpdatelikeRecordMutationError = SchemaErrorResponse

    /**
 * @summary 定期的に実行される。1h経過したらlike数を更新する
 */
export const useUpdatelikeRecord = <TError = SchemaErrorResponse,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof updatelikeRecord>>, TError,void, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationResult<
        Awaited<ReturnType<typeof updatelikeRecord>>,
        TError,
        void,
        TContext
      > => {

      const mutationOptions = getUpdatelikeRecordMutationOptions(options);

      return useMutation(mutationOptions);
    }
    
/**
 * @summary 各投稿の盛り上がり度を取得
 */
export const listSpeeds = (
    
 options?: SecondParameter<typeof customInstance>,signal?: AbortSignal
) => {
      
      
      return customInstance<SchemaSpeedResponse[]>(
      {url: `/posts/speed`, method: 'GET', signal
    },
      options);
    }
  

export const getListSpeedsQueryKey = () => {
    return [`/posts/speed`] as const;
    }

    
export const getListSpeedsQueryOptions = <TData = Awaited<ReturnType<typeof listSpeeds>>, TError = SchemaErrorResponse>( options?: { query?:Partial<UseQueryOptions<Awaited<ReturnType<typeof listSpeeds>>, TError, TData>>, request?: SecondParameter<typeof customInstance>}
) => {

const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey =  queryOptions?.queryKey ?? getListSpeedsQueryKey();

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof listSpeeds>>> = ({ signal }) => listSpeeds(requestOptions, signal);

      

      

   return  { queryKey, queryFn, ...queryOptions} as UseQueryOptions<Awaited<ReturnType<typeof listSpeeds>>, TError, TData> & { queryKey: QueryKey }
}

export type ListSpeedsQueryResult = NonNullable<Awaited<ReturnType<typeof listSpeeds>>>
export type ListSpeedsQueryError = SchemaErrorResponse

/**
 * @summary 各投稿の盛り上がり度を取得
 */
export const useListSpeeds = <TData = Awaited<ReturnType<typeof listSpeeds>>, TError = SchemaErrorResponse>(
  options?: { query?:Partial<UseQueryOptions<Awaited<ReturnType<typeof listSpeeds>>, TError, TData>>, request?: SecondParameter<typeof customInstance>}

  ):  UseQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getListSpeedsQueryOptions(options)

  const query = useQuery(queryOptions) as  UseQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryOptions.queryKey ;

  return query;
}




